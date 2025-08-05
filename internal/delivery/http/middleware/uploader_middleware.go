package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type UploadOptions struct {
	FieldName     string
	MaxFileSizeMB int64
	MaxFiles      int
	BucketName    string
	AllowedTypes  []string
}

func SingleFileUpload(minioClient *minio.Client, opts UploadOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := GetUser(c)
		file, err := c.FormFile(opts.FieldName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Field '%s' tidak ditemukan", opts.FieldName)})
			return
		}

		if file.Size > opts.MaxFileSizeMB*1024*1024 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Ukuran file terlalu besar"})
			return
		}

		src, err := file.Open()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuka file"})
			return
		}
		defer src.Close()

		buffer := make([]byte, 512)
		if _, err := src.Read(buffer); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca file"})
			return
		}
		contentType := http.DetectContentType(buffer)

		if _, err := src.Seek(0, io.SeekStart); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal reset posisi file"})
			return
		}

		if len(opts.AllowedTypes) > 0 {
			allowed := slices.Contains(opts.AllowedTypes, contentType)
			if !allowed {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Tipe file tidak diizinkan: %s", contentType)})
				return
			}
		}

		now := time.Now()
		ext := filepath.Ext(file.Filename)
		formattedTime := now.Format("02-01-2006-15-04-05-000")
		uniqueName := fmt.Sprintf("%v-%s%s", auth.ID, formattedTime, ext)

		metadata := map[string]string{
			"original_name": file.Filename,
			"author_id":     auth.ID.String(),
			"created_at":    now.Format(time.RFC3339Nano),
			"updated_at":    now.Format(time.RFC3339Nano),
			"file_size":     fmt.Sprintf("%d", file.Size),
		}

		info, err := minioClient.PutObject(
			context.Background(),
			opts.BucketName,
			uniqueName,
			src,
			file.Size,
			minio.PutObjectOptions{
				ContentType:  contentType,
				UserMetadata: metadata,
			},
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal upload ke MinIO"})
			return
		}

		c.Set("uploadedFile", map[string]any{
			"file_name":     uniqueName,
			"original_name": file.Filename,
			"file_size":     file.Size,
			"url":           fmt.Sprintf("/%s/%s", info.Bucket, info.Key),
			"author_id":     auth.ID,
			"created_at":    now,
			"updated_at":    now,
		})

		c.Next()
	}
}

func MultipleFileUpload(minioClient *minio.Client, opts UploadOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := GetUser(c)

		form, err := c.MultipartForm()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Gagal membaca form data"})
			return
		}

		files := form.File[opts.FieldName]
		if len(files) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Field '%s' tidak ditemukan", opts.FieldName)})
			return
		}

		var uploadedFiles []map[string]any

		for _, file := range files {
			if file.Size > opts.MaxFileSizeMB*1024*1024 {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ukuran file '%s' terlalu besar", file.Filename)})
				return
			}

			src, err := file.Open()
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal membuka file '%s'", file.Filename)})
				return
			}

			buffer := make([]byte, 512)
			if _, err := src.Read(buffer); err != nil {
				src.Close()
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca file"})
				return
			}
			contentType := http.DetectContentType(buffer)

			if _, err := src.Seek(0, io.SeekStart); err != nil {
				src.Close()
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal reset posisi file"})
				return
			}

			if len(opts.AllowedTypes) > 0 {
				allowed := slices.Contains(opts.AllowedTypes, contentType)
				if !allowed {
					src.Close()
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Tipe file tidak diizinkan: %s", contentType)})
					return
				}
			}

			now := time.Now()
			ext := filepath.Ext(file.Filename)
			formattedTime := now.Format("02-01-2006-15-04-05-000")
			uniqueName := fmt.Sprintf("%v-%s%s", auth.ID, formattedTime, ext)

			metadata := map[string]string{
				"original_name": file.Filename,
				"author_id":     auth.ID.String(),
				"created_at":    now.Format(time.RFC3339Nano),
				"updated_at":    now.Format(time.RFC3339Nano),
				"file_size":     fmt.Sprintf("%d", file.Size),
			}

			info, err := minioClient.PutObject(
				context.Background(),
				opts.BucketName,
				uniqueName,
				src,
				file.Size,
				minio.PutObjectOptions{
					ContentType:  contentType,
					UserMetadata: metadata,
				},
			)
			src.Close()

			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal upload file '%s'", file.Filename)})
				return
			}

			uploadedFiles = append(uploadedFiles, map[string]any{
				"file_name":     uniqueName,
				"original_name": file.Filename,
				"file_size":     file.Size,
				"url":           fmt.Sprintf("/%s/%s", info.Bucket, info.Key),
				"author_id":     auth.ID,
				"created_at":    now,
				"updated_at":    now,
			})
		}

		c.Set("uploadedFiles", uploadedFiles)
		c.Next()
	}
}

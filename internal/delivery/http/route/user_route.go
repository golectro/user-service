package route

import (
	"golectro-user/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (c *RouteConfig) RegisterUserRoutes(rg *gin.RouterGroup, minioClient *minio.Client) {
	user := rg.Group("/users")

	user.POST("/sync", c.AuthMiddleware, c.UserController.SyncUser)
	user.POST("/avatar", c.AuthMiddleware, middleware.SingleFileUpload(minioClient, middleware.UploadOptions{
		FieldName:     "file",
		MaxFileSizeMB: 5,
		BucketName:    c.Viper.GetString("MINIO_BUCKET_AVATAR"),
		AllowedTypes:  []string{"image/jpeg", "image/png", "image/gif"},
	}), c.UserController.UploadAvatar)

	user.GET("/avatar/download", c.AuthMiddleware, c.UserController.DownloadAvatar)
	user.GET("/avatar/preview", c.AuthMiddleware, c.UserController.PreviewAvatar)
}

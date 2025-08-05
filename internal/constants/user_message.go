package constants

import "golectro-user/internal/model"

var (
	UserSynced = model.Message{
		"en": "User synced successfully",
		"id": "Pengguna berhasil disinkronkan",
	}
	AvatarUploaded = model.Message{
		"en": "Avatar uploaded successfully",
		"id": "Avatar berhasil diunggah",
	}
	AvatarDownloaded = model.Message{
		"en": "Avatar downloaded successfully",
		"id": "Avatar berhasil diunduh",
	}
)

var (
	FailedSyncUser = model.Message{
		"en": "Failed to sync user",
		"id": "Gagal menyinkronkan pengguna",
	}
	FailedUploadAvatar = model.Message{
		"en": "Failed to upload avatar",
		"id": "Gagal mengunggah avatar",
	}
	InvalidFileType = model.Message{
		"en": "Invalid file type",
		"id": "Tipe file tidak valid",
	}
	InvalidResetPosition = model.Message{
		"en": "Failed to reset file position",
		"id": "Gagal mengatur ulang posisi file",
	}
	InvalidReadFile = model.Message{
		"en": "Failed to read file",
		"id": "Gagal membaca file",
	}
	InvalidOpenFile = model.Message{
		"en": "Failed to open file",
		"id": "Gagal membuka file",
	}
	FileSizeExceeded = model.Message{
		"en": "File size exceeded",
		"id": "Ukuran file melebihi batas",
	}
	FileNotFound = model.Message{
		"en": "File not found",
		"id": "File tidak ditemukan",
	}
	FailedGetPresignedURL = model.Message{
		"en": "Failed to get presigned URL",
		"id": "Gagal mendapatkan URL presigned",
	}
	FailedFindUserByID = model.Message{
		"en": "Failed to find user by ID",
		"id": "Gagal menemukan pengguna berdasarkan ID",
	}
	UserNotFound = model.Message{
		"en": "User not found",
		"id": "Pengguna tidak ditemukan",
	}
)

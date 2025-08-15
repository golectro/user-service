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
	FailedUploadAvatar = model.Message{
		"en": "Failed to upload avatar",
		"id": "Gagal mengunggah avatar",
	}
)

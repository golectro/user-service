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
)

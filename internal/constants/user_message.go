package constants

import "golectro-user/internal/model"

var (
	UserSynced = model.Message{
		"en": "User synced successfully",
		"id": "Pengguna berhasil disinkronkan",
	}
)

var (
	FailedSyncUser = model.Message{
		"en": "Failed to sync user",
		"id": "Gagal menyinkronkan pengguna",
	}
)

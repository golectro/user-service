package constants

import "golectro-user/internal/model"

var (
	AddressesRetrieved = model.Message{
		"en": "Addresses retrieved successfully",
		"id": "Alamat berhasil diambil",
	}
)

var (
	FailedGetAddresses = model.Message{
		"en": "Failed to get addresses",
		"id": "Gagal mendapatkan alamat",
	}
)

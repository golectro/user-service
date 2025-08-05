package constants

import "golectro-user/internal/model"

var (
	AddressesRetrieved = model.Message{
		"en": "Addresses retrieved successfully",
		"id": "Alamat berhasil diambil",
	}
	UpdateAddressSuccess = model.Message{
		"en": "Address updated successfully",
		"id": "Alamat berhasil diperbarui",
	}
	AddressCreated = model.Message{
		"en": "Address created successfully",
		"id": "Alamat berhasil dibuat",
	}
)

var (
	FailedGetAddresses = model.Message{
		"en": "Failed to get addresses",
		"id": "Gagal mendapatkan alamat",
	}
	FailedCreateAddress = model.Message{
		"en": "Failed to create address",
		"id": "Gagal membuat alamat",
	}
	FailedUpdateAddress = model.Message{
		"en": "Failed to update address",
		"id": "Gagal memperbarui alamat",
	}
)

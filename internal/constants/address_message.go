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
	SetDefaultAddressSuccess = model.Message{
		"en": "Default address set successfully",
		"id": "Alamat default berhasil diatur",
	}
	AddressDeleted = model.Message{
		"en": "Address deleted successfully",
		"id": "Alamat berhasil dihapus",
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
	AddressNotFound = model.Message{
		"en": "Address not found",
		"id": "Alamat tidak ditemukan",
	}
	FailedDeleteAddress = model.Message{
		"en": "Failed to delete address",
		"id": "Gagal menghapus alamat",
	}
	FailedSetDefaultAddress = model.Message{
		"en": "Failed to set default address",
		"id": "Gagal mengatur alamat default",
	}
)

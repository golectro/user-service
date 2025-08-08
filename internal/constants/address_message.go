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
	FailedGeneateDEK = model.Message{
		"en": "Failed to generate DEK",
		"id": "Gagal membuat kunci enkripsi",
	}
	FailedEncryptLabel = model.Message{
		"en": "Failed to encrypt label",
		"id": "Gagal enkripsi label",
	}
	FailedEncryptRecipient = model.Message{
		"en": "Failed to encrypt recipient",
		"id": "Gagal enkripsi penerima",
	}
	FailedEncryptPhone = model.Message{
		"en": "Failed to encrypt phone",
		"id": "Gagal enkripsi telepon",
	}
	FailedEncryptAddressLine = model.Message{
		"en": "Failed to encrypt address line",
		"id": "Gagal enkripsi baris alamat",
	}
	FailedEncryptCity = model.Message{
		"en": "Failed to encrypt city",
		"id": "Gagal enkripsi kota",
	}
	FailedEncryptProvince = model.Message{
		"en": "Failed to encrypt province",
		"id": "Gagal enkripsi provinsi",
	}
	FailedEncryptPostalCode = model.Message{
		"en": "Failed to encrypt postal code",
		"id": "Gagal enkripsi kode pos",
	}
	FailedEncryptDEK = model.Message{
		"en": "Failed to encrypt DEK",
		"id": "Gagal enkripsi DEK",
	}
)

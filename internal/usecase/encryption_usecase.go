package usecase

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golectro-user/internal/constants"
	"golectro-user/internal/model"
	"golectro-user/internal/model/converter"
	"golectro-user/internal/repository"
	"golectro-user/internal/utils"
	"io"

	"github.com/google/uuid"
	vault "github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type EncryptionUsecase struct {
	DB                   *gorm.DB
	Log                  *logrus.Logger
	Client               *vault.Client
	Viper                *viper.Viper
	EncryptionRepository *repository.EncryptionRepository
}

func NewEncryptionUsecase(db *gorm.DB, log *logrus.Logger, client *vault.Client, viper *viper.Viper, encryptionRepository *repository.EncryptionRepository) *EncryptionUsecase {
	return &EncryptionUsecase{
		DB:                   db,
		Log:                  log,
		Client:               client,
		Viper:                viper,
		EncryptionRepository: encryptionRepository,
	}
}

func (uc *EncryptionUsecase) GenerateDEK() ([]byte, error) {
	dek := make([]byte, 32)
	if _, err := rand.Read(dek); err != nil {
		return nil, fmt.Errorf("failed to generate DEK: %w", err)
	}
	return dek, nil
}

func (uc *EncryptionUsecase) EncryptDEK(dek []byte) (string, error) {
	plaintext := base64.StdEncoding.EncodeToString(dek)

	fmt.Println("ENV transit key:", uc.Viper.GetString("VAULT_TRANSIT_KEY"))

	secret, err := uc.Client.Logical().Write(fmt.Sprintf("transit/encrypt/%s", uc.Viper.GetString("VAULT_TRANSIT_KEY")), map[string]interface{}{
		"plaintext": plaintext,
	})
	if err != nil {
		return "", err
	}
	return secret.Data["ciphertext"].(string), nil
}

func (uc *EncryptionUsecase) DecryptDEK(ciphertext string) ([]byte, error) {
	secret, err := uc.Client.Logical().Write(fmt.Sprintf("transit/decrypt/%s", uc.Viper.GetString("VAULT_TRANSIT_KEY")), map[string]interface{}{
		"ciphertext": ciphertext,
	})
	if err != nil {
		return nil, err
	}
	b64 := secret.Data["plaintext"].(string)
	return base64.StdEncoding.DecodeString(b64)
}

func (uc *EncryptionUsecase) EncryptAES_GCM(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	return append(nonce, ciphertext...), nil
}

func (uc *EncryptionUsecase) DecryptAES_GCM(ciphertextWithNonce, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertextWithNonce) < 12 {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce := ciphertextWithNonce[:12]
	ciphertext := ciphertextWithNonce[12:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func (uc *EncryptionUsecase) GetAddressEncryptionKey(ctx context.Context, userID, addressID uuid.UUID) (*model.AddressEncryptionKeyResponse, error) {
	tx := uc.DB.WithContext(ctx)

	keyEntity, err := uc.EncryptionRepository.FindByAddressIDAndUserID(tx, userID, addressID)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to find encryption key by address ID")
		return nil, utils.WrapMessageAsError(constants.FailedGetEncryptionKey, err)
	}

	if keyEntity == nil {
		uc.Log.Error("Encryption key not found for address ID")
		return nil, utils.WrapMessageAsError(constants.EncryptionKeyNotFound, nil)
	}

	return converter.EncryptionKeyToResponse(keyEntity), nil
}

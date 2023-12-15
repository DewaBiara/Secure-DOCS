package impl

import (
	"context"

	"github.com/DewaBiara/Secure-DOCS/internal/secure/dto"
	"github.com/DewaBiara/Secure-DOCS/internal/secure/repository"
	"github.com/DewaBiara/Secure-DOCS/internal/secure/service"
	"github.com/DewaBiara/Secure-DOCS/pkg/utils/aes"
	"github.com/google/uuid"
)

type (
	EncryptionServiceImpl struct {
		encryptionRepository repository.EncryptionRepository
		crypter              aes.AESFileCrypter
	}
)

func NewEncryptionServiceImpl(encryptionRepository repository.EncryptionRepository, crypter aes.AESFileCrypter) service.EncryptionService {
	return &EncryptionServiceImpl{
		encryptionRepository: encryptionRepository,
		crypter:              crypter,
	}
}

func (u *EncryptionServiceImpl) CreateEncryption(ctx context.Context, Encryption *dto.CreateEncryptionRequest) error {

	EncryptionEntity := Encryption.ToEntity()
	EncryptionEntity.ID = uint(uuid.New().ID())

	err := u.encryptionRepository.CreateEncryption(ctx, EncryptionEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *EncryptionServiceImpl) GetSingleEncryption(ctx context.Context, encryptionID string) (*dto.GetSingleEncryptionResponse, error) {
	encryption, err := d.encryptionRepository.GetSingleEncryption(ctx, encryptionID)
	if err != nil {
		return nil, err
	}

	var encryptionResponse = dto.NewGetSingleEncryptionResponse(encryption)

	return encryptionResponse, nil
}

func (u *EncryptionServiceImpl) GetPageEncryption(ctx context.Context, page int, limit int) (*dto.GetPageEncryptionsResponse, error) {
	offset := (page - 1) * limit

	encryptions, err := u.encryptionRepository.GetPageEncryption(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageEncryptionsResponse(encryptions), nil
}

func (u *EncryptionServiceImpl) UpdateEncryption(ctx context.Context, encryptionID uint, updateEncryption *dto.UpdateEncryptionRequest) error {
	encryption := updateEncryption.ToEntity()
	encryption.ID = encryptionID

	return u.encryptionRepository.UpdateEncryption(ctx, encryption)
}

func (d *EncryptionServiceImpl) DeleteEncryption(ctx context.Context, encryptionID string) error {

	return d.encryptionRepository.DeleteEncryption(ctx, encryptionID)
}

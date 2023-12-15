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
	DecryptionServiceImpl struct {
		decryptionRepository repository.DecryptionRepository
		fileCrypter          aes.FileCrypter
	}
)

func NewDecryptionServiceImpl(decryptionRepository repository.DecryptionRepository, fileCrypter aes.FileCrypter) service.DecryptionService {
	return &DecryptionServiceImpl{
		decryptionRepository: decryptionRepository,
		fileCrypter:          fileCrypter,
	}
}

func (u *DecryptionServiceImpl) CreateDecryption(ctx context.Context, Decryption *dto.CreateDecryptionRequest) error {

	DecryptionEntity := Decryption.ToEntity()
	DecryptionEntity.ID = uint(uuid.New().ID())

	err := u.decryptionRepository.CreateDecryption(ctx, DecryptionEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *DecryptionServiceImpl) GetSingleDecryption(ctx context.Context, decryptionID string) (*dto.GetSingleDecryptionResponse, error) {
	decryption, err := d.decryptionRepository.GetSingleDecryption(ctx, decryptionID)
	if err != nil {
		return nil, err
	}

	var decryptionResponse = dto.NewGetSingleDecryptionResponse(decryption)

	return decryptionResponse, nil
}

func (u *DecryptionServiceImpl) GetPageDecryption(ctx context.Context, page int, limit int) (*dto.GetPageDecryptionsResponse, error) {
	offset := (page - 1) * limit

	decryptions, err := u.decryptionRepository.GetPageDecryption(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageDecryptionsResponse(decryptions), nil
}

func (u *DecryptionServiceImpl) UpdateDecryption(ctx context.Context, decryptionID uint, updateDecryption *dto.UpdateDecryptionRequest) error {
	decryption := updateDecryption.ToEntity()
	decryption.ID = decryptionID

	return u.decryptionRepository.UpdateDecryption(ctx, decryption)
}

func (d *DecryptionServiceImpl) DeleteDecryption(ctx context.Context, decryptionID string) error {

	return d.decryptionRepository.DeleteDecryption(ctx, decryptionID)
}

package impl

import (
	"context"

	"github.com/DewaBiara/Secure-DOCS/internal/secure/dto"
	"github.com/DewaBiara/Secure-DOCS/internal/secure/repository"
	"github.com/DewaBiara/Secure-DOCS/internal/secure/service"
	"github.com/google/uuid"
)

type (
	KeyServiceImpl struct {
		keyRepository repository.KeyRepository
	}
)

func NewKeyServiceImpl(keyRepository repository.KeyRepository) service.KeyService {
	return &KeyServiceImpl{
		keyRepository: keyRepository,
	}
}

func (u *KeyServiceImpl) CreateKey(ctx context.Context, key *dto.CreateKeyRequest) error {

	keyEntity := key.ToEntity()
	keyEntity.ID = uint(uuid.New().ID())

	err := u.keyRepository.CreateKey(ctx, keyEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *KeyServiceImpl) GetSingleKey(ctx context.Context, keyID string) (*dto.GetSingleKeyResponse, error) {
	key, err := d.keyRepository.GetSingleKey(ctx, keyID)
	if err != nil {
		return nil, err
	}

	var keyResponse = dto.NewGetSingleKeyResponse(key)

	return keyResponse, nil
}

func (u *KeyServiceImpl) GetPageKey(ctx context.Context, page int, limit int) (*dto.GetPageKeysResponse, error) {
	offset := (page - 1) * limit

	keys, err := u.keyRepository.GetPageKey(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return dto.NewGetPageKeysResponse(keys), nil
}

func (u *KeyServiceImpl) UpdateKey(ctx context.Context, keyID uint, updateKey *dto.UpdateKeyRequest) error {
	key := updateKey.ToEntity()
	key.ID = keyID

	return u.keyRepository.UpdateKey(ctx, key)
}

func (d *KeyServiceImpl) DeleteKey(ctx context.Context, keyID string) error {

	return d.keyRepository.DeleteKey(ctx, keyID)
}

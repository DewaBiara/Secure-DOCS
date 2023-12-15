package impl

import (
	"context"
	"strings"

	"github.com/DewaBiara/Secure-DOCS/internal/secure/repository"
	"github.com/DewaBiara/Secure-DOCS/pkg/entity"
	"github.com/DewaBiara/Secure-DOCS/pkg/utils"
	"gorm.io/gorm"
)

type KeyRepositoryImpl struct {
	db *gorm.DB
}

func NewKeyRepositoryImpl(db *gorm.DB) repository.KeyRepository {
	keyRepository := &KeyRepositoryImpl{
		db: db,
	}

	return keyRepository
}

func (u *KeyRepositoryImpl) CreateKey(ctx context.Context, key *entity.Key) error {
	err := u.db.WithContext(ctx).Create(key).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(err.Error(), "name"):
				return utils.ErrKeyAlreadyExist
			}
		}

		return err
	}

	return nil
}

func (u *KeyRepositoryImpl) GetSingleKey(ctx context.Context, keyID string) (*entity.Key, error) {
	var key entity.Key
	err := u.db.WithContext(ctx).Select([]string{"id", "pengirim_id", "penerima_id", "encryption_id", "key"}).
		Where("id = ?", keyID).First(&key).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrKeyNotFound
		}

		return nil, err
	}

	return &key, nil
}

func (u *KeyRepositoryImpl) GetPageKey(ctx context.Context, limit int, offset int) (*entity.Keys, error) {
	var keys entity.Keys
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&keys).Error
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, utils.ErrKeyNotFound
	}

	return &keys, nil
}

func (u *KeyRepositoryImpl) GetPageKeyByPenerima(ctx context.Context, penerimaID string, limit int, offset int) (*entity.Keys, error) {
	var keys entity.Keys
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Where("penerima_id = ?", penerimaID).
		Offset(offset).
		Limit(limit).
		Find(&keys).Error
	if err != nil {
		return nil, err
	}

	if len(keys) == 0 {
		return nil, utils.ErrKeyNotFound
	}

	return &keys, nil
}

func (u *KeyRepositoryImpl) UpdateKey(ctx context.Context, key *entity.Key) error {
	result := u.db.WithContext(ctx).Model(&entity.Key{}).Where("id = ?", key.ID).Updates(key)
	if result.Error != nil {
		errStr := result.Error.Error()
		if strings.Contains(errStr, "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(errStr, "name"):
				return utils.ErrKeyAlreadyExist
			}
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrKeyNotFound
	}

	return nil
}

func (d *KeyRepositoryImpl) DeleteKey(ctx context.Context, keyID string) error {
	result := d.db.WithContext(ctx).
		Select("Key").
		Delete(&entity.Key{}, "id = ?", keyID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrKeyNotFound
	}

	return nil
}

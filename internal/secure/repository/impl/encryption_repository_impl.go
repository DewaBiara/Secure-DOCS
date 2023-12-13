package impl

import (
	"context"
	"strings"

	"github.com/DewaBiara/Secure-DOCS/internal/secure/repository"
	"github.com/DewaBiara/Secure-DOCS/pkg/entity"
	"github.com/DewaBiara/Secure-DOCS/pkg/utils"
	"gorm.io/gorm"
)

type EncryptionRepositoryImpl struct {
	db *gorm.DB
}

func NewEncryptionRepositoryImpl(db *gorm.DB) repository.EncryptionRepository {
	encryptionRepository := &EncryptionRepositoryImpl{
		db: db,
	}

	return encryptionRepository
}

func (u *EncryptionRepositoryImpl) CreateEncryption(ctx context.Context, encryption *entity.Encryption) error {
	err := u.db.WithContext(ctx).Create(encryption).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(err.Error(), "name"):
				return utils.ErrEncryptionAlreadyExist
			}
		}

		return err
	}

	return nil
}

func (u *EncryptionRepositoryImpl) GetSingleEncryption(ctx context.Context, encryptionID string) (*entity.Encryption, error) {
	var encryption entity.Encryption
	err := u.db.WithContext(ctx).Select([]string{"id", "userid", "keyid", "filename", "status"}).
		Where("id = ?", encryptionID).First(&encryption).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrEncryptionNotFound
		}

		return nil, err
	}

	return &encryption, nil
}

func (u *EncryptionRepositoryImpl) GetPageEncryption(ctx context.Context, limit int, offset int) (*entity.Encryptions, error) {
	var encryptions entity.Encryptions
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&encryptions).Error
	if err != nil {
		return nil, err
	}

	if len(encryptions) == 0 {
		return nil, utils.ErrEncryptionNotFound
	}

	return &encryptions, nil
}

func (u *EncryptionRepositoryImpl) UpdateEncryption(ctx context.Context, encryption *entity.Encryption) error {
	result := u.db.WithContext(ctx).Model(&entity.Encryption{}).Where("id = ?", encryption.ID).Updates(encryption)
	if result.Error != nil {
		errStr := result.Error.Error()
		if strings.Contains(errStr, "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(errStr, "name"):
				return utils.ErrEncryptionAlreadyExist
			}
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrEncryptionNotFound
	}

	return nil
}

func (d *EncryptionRepositoryImpl) DeleteEncryption(ctx context.Context, encryptionID string) error {
	result := d.db.WithContext(ctx).
		Select("Encryption").
		Delete(&entity.Encryption{}, "id = ?", encryptionID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrEncryptionNotFound
	}

	return nil
}

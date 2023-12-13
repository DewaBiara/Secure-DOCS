package impl

import (
	"context"
	"strings"

	"github.com/DewaBiara/Secure-DOCS/internal/secure/repository"
	"github.com/DewaBiara/Secure-DOCS/pkg/entity"
	"github.com/DewaBiara/Secure-DOCS/pkg/utils"
	"gorm.io/gorm"
)

type DecryptionRepositoryImpl struct {
	db *gorm.DB
}

func NewDecryptionRepositoryImpl(db *gorm.DB) repository.DecryptionRepository {
	decryptionRepository := &DecryptionRepositoryImpl{
		db: db,
	}

	return decryptionRepository
}

func (u *DecryptionRepositoryImpl) CreateDecryption(ctx context.Context, decryption *entity.Decryption) error {
	err := u.db.WithContext(ctx).Create(decryption).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(err.Error(), "name"):
				return utils.ErrDecryptionAlreadyExist
			}
		}

		return err
	}

	return nil
}

func (u *DecryptionRepositoryImpl) GetSingleDecryption(ctx context.Context, decryptionID string) (*entity.Decryption, error) {
	var decryption entity.Decryption
	err := u.db.WithContext(ctx).Select([]string{"id", ""}).
		Where("id = ?", decryptionID).First(&decryption).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrDecryptionNotFound
		}

		return nil, err
	}

	return &decryption, nil
}

func (u *DecryptionRepositoryImpl) GetPageDecryption(ctx context.Context, limit int, offset int) (*entity.Decryptions, error) {
	var decryptions entity.Decryptions
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&decryptions).Error
	if err != nil {
		return nil, err
	}

	if len(decryptions) == 0 {
		return nil, utils.ErrDecryptionNotFound
	}

	return &decryptions, nil
}

func (u *DecryptionRepositoryImpl) UpdateDecryption(ctx context.Context, decryption *entity.Decryption) error {
	result := u.db.WithContext(ctx).Model(&entity.Decryption{}).Where("id = ?", decryption.ID).Updates(decryption)
	if result.Error != nil {
		errStr := result.Error.Error()
		if strings.Contains(errStr, "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(errStr, "name"):
				return utils.ErrDecryptionAlreadyExist
			}
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrDecryptionNotFound
	}

	return nil
}

func (d *DecryptionRepositoryImpl) DeleteDecryption(ctx context.Context, decryptionID string) error {
	result := d.db.WithContext(ctx).
		Select("Decryption").
		Delete(&entity.Decryption{}, "id = ?", decryptionID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrDecryptionNotFound
	}

	return nil
}

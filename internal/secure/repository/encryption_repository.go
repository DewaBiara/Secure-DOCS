package repository

import (
	"context"

	"github.com/DewaBiara/Secure-DOCS/pkg/entity"
)

type EncryptionRepository interface {
	CreateEncryption(ctx context.Context, encryption *entity.Encryption) error
	UpdateEncryption(ctx context.Context, encryption *entity.Encryption) error
	GetSingleEncryption(ctx context.Context, encryptionID string) (*entity.Encryption, error)
	GetPageEncryption(ctx context.Context, limit int, offset int) (*entity.Encryptions, error)
	DeleteEncryption(ctx context.Context, encryptionID string) error
}

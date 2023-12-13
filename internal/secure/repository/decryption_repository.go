package repository

import (
	"context"

	"github.com/DewaBiara/Secure-DOCS/pkg/entity"
)

type DecryptionRepository interface {
	CreateDecryption(ctx context.Context, decryption *entity.Decryption) error
	UpdateDecryption(ctx context.Context, decryption *entity.Decryption) error
	GetSingleDecryption(ctx context.Context, decryptionID string) (*entity.Decryption, error)
	GetPageDecryption(ctx context.Context, limit int, offset int) (*entity.Decryptions, error)
	DeleteDecryption(ctx context.Context, decryptionID string) error
}

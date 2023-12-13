package service

import (
	"context"

	"github.com/DewaBiara/Secure-DOCS/internal/secure/dto"
)

type DecryptionService interface {
	CreateDecryption(ctx context.Context, decryption *dto.CreateDecryptionRequest) error
	UpdateDecryption(ctx context.Context, decryptionID uint, updateDecryption *dto.UpdateDecryptionRequest) error
	GetSingleDecryption(ctx context.Context, decryptionID string) (*dto.GetSingleDecryptionResponse, error)
	GetPageDecryption(ctx context.Context, limit int, offset int) (*dto.GetPageDecryptionsResponse, error)
	DeleteDecryption(ctx context.Context, decryptionID string) error
}

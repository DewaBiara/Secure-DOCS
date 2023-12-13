package service

import (
	"context"

	"github.com/DewaBiara/Secure-DOCS/internal/secure/dto"
)

type EncryptionService interface {
	CreateEncryption(ctx context.Context, encryption *dto.CreateEncryptionRequest) error
	UpdateEncryption(ctx context.Context, encryptionID uint, updateEncryption *dto.UpdateEncryptionRequest) error
	GetSingleEncryption(ctx context.Context, encryptionID string) (*dto.GetSingleEncryptionResponse, error)
	GetPageEncryption(ctx context.Context, limit int, offset int) (*dto.GetPageEncryptionsResponse, error)
	DeleteEncryption(ctx context.Context, encryptionID string) error
}

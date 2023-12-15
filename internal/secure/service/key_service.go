package service

import (
	"context"

	"github.com/DewaBiara/Secure-DOCS/internal/secure/dto"
)

type KeyService interface {
	CreateKey(ctx context.Context, key *dto.CreateKeyRequest) error
	UpdateKey(ctx context.Context, keyID uint, updateKey *dto.UpdateKeyRequest) error
	GetSingleKey(ctx context.Context, keyID string) (*dto.GetSingleKeyResponse, error)
	GetPageKey(ctx context.Context, limit int, offset int) (*dto.GetPageKeysResponse, error)
	GetPageKeyByPenerima(ctx context.Context, penerimaID string, limit int, offset int) (*dto.GetPageKeysResponse, error)
	DeleteKey(ctx context.Context, keyID string) error
}

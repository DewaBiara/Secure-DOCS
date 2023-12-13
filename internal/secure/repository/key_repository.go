package repository

import (
	"context"

	"github.com/DewaBiara/Secure-DOCS/pkg/entity"
)

type KeyRepository interface {
	CreateKey(ctx context.Context, key *entity.Key) error
	UpdateKey(ctx context.Context, key *entity.Key) error
	GetSingleKey(ctx context.Context, keyID string) (*entity.Key, error)
	GetPageKey(ctx context.Context, limit int, offset int) (*entity.Keys, error)
	DeleteKey(ctx context.Context, keyID string) error
}

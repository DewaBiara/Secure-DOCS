package dto

import (
	"github.com/DewaBiara/Secure-DOCS/pkg/entity"
)

type CreateKeyRequest struct {
	PengirimID   string `json:"pengirimid" validate:"required"`
	PenerimaID   string `json:"penerimaid" validate:"required"`
	EncryptionID uint   `json:"encryptionid" validate:"required"`
	Key          string `json:"key" validate:"required"`
}

func (u *CreateKeyRequest) ToEntity() *entity.Key {
	return &entity.Key{
		PengirimID:   u.PengirimID,
		PenerimaID:   u.PenerimaID,
		EncryptionID: u.EncryptionID,
		Key:          u.Key,
	}
}

type UpdateKeyRequest struct {
	ID           uint   `json:"id" validate:"required"`
	PengirimID   string `json:"pengirimid" validate:"required"`
	PenerimaID   string `json:"penerimaid" validate:"required"`
	EncryptionID uint   `json:"encryptionid" validate:"required"`
	Key          string `json:"key" validate:"required"`
}

func (u *UpdateKeyRequest) ToEntity() *entity.Key {
	return &entity.Key{
		PengirimID:   u.PengirimID,
		PenerimaID:   u.PenerimaID,
		EncryptionID: u.EncryptionID,
		Key:          u.Key,
	}
}

type GetSingleKeyResponse struct {
	ID           uint   `json:"id"`
	PengirimID   string `json:"pengirimid" validate:"required"`
	PenerimaID   string `json:"penerimaid" validate:"required"`
	EncryptionID uint   `json:"encryptionid" validate:"required"`
	Key          string `json:"key" validate:"required"`
}

func NewGetSingleKeyResponse(key *entity.Key) *GetSingleKeyResponse {
	return &GetSingleKeyResponse{
		ID:           key.ID,
		PengirimID:   key.PengirimID,
		PenerimaID:   key.PenerimaID,
		EncryptionID: key.EncryptionID,
		Key:          key.Key,
	}
}

type GetPageKeyResponse struct {
	ID           uint   `json:"id"`
	PengirimID   string `json:"pengirimid" validate:"required"`
	PenerimaID   string `json:"penerimaid" validate:"required"`
	EncryptionID uint   `json:"encryptionid" validate:"required"`
	Key          string `json:"key" validate:"required"`
}

func NewGetPageKeyResponse(key *entity.Key) *GetPageKeyResponse {
	return &GetPageKeyResponse{
		ID:           key.ID,
		PengirimID:   key.PengirimID,
		PenerimaID:   key.PenerimaID,
		EncryptionID: key.EncryptionID,
		Key:          key.Key,
	}
}

type GetPageKeysResponse []GetPageKeyResponse

func NewGetPageKeysResponse(keys *entity.Keys) *GetPageKeysResponse {
	var getPageKeys GetPageKeysResponse
	for _, keys := range *keys {
		getPageKeys = append(getPageKeys, *NewGetPageKeyResponse(&keys))
	}
	return &getPageKeys
}

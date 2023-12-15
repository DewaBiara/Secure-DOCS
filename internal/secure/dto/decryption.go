package dto

import (
	"github.com/DewaBiara/Secure-DOCS/pkg/entity"
)

type CreateDecryptionRequest struct {
	UserID   string `json:"userid" validate:"required"`
	KeyID    uint   `json:"keyid" validate:"required"`
	Filename string `json:"filename" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

func (u *CreateDecryptionRequest) ToEntity() *entity.Decryption {
	return &entity.Decryption{
		UserID:   u.UserID,
		KeyID:    u.KeyID,
		Filename: u.Filename,
		Status:   u.Status,
	}
}

type UpdateDecryptionRequest struct {
	ID       uint   `json:"id" validate:"required"`
	UserID   string `json:"userid" validate:"required"`
	KeyID    uint   `json:"keyid" validate:"required"`
	Filename string `json:"filename" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

func (u *UpdateDecryptionRequest) ToEntity() *entity.Decryption {
	return &entity.Decryption{
		UserID:   u.UserID,
		KeyID:    u.KeyID,
		Filename: u.Filename,
		Status:   u.Status,
	}
}

type GetSingleDecryptionResponse struct {
	ID       uint   `json:"id"`
	UserID   string `json:"userid" validate:"required"`
	KeyID    uint   `json:"keyid" validate:"required"`
	Filename string `json:"filename" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

func NewGetSingleDecryptionResponse(decryption *entity.Decryption) *GetSingleDecryptionResponse {
	return &GetSingleDecryptionResponse{
		ID:       decryption.ID,
		UserID:   decryption.UserID,
		KeyID:    decryption.KeyID,
		Filename: decryption.Filename,
		Status:   decryption.Status,
	}
}

type GetPageDecryptionResponse struct {
	ID       uint   `json:"id"`
	UserID   string `json:"userid" validate:"required"`
	KeyID    uint   `json:"keyid" validate:"required"`
	Filename string `json:"filename" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

func NewGetPageDecryptionResponse(decryption *entity.Decryption) *GetPageDecryptionResponse {
	return &GetPageDecryptionResponse{
		ID:       decryption.ID,
		UserID:   decryption.UserID,
		KeyID:    decryption.KeyID,
		Filename: decryption.Filename,
		Status:   decryption.Status,
	}
}

type GetPageDecryptionsResponse []GetPageDecryptionResponse

func NewGetPageDecryptionsResponse(decryptions *entity.Decryptions) *GetPageDecryptionsResponse {
	var getPageDecryptions GetPageDecryptionsResponse
	for _, decryptions := range *decryptions {
		getPageDecryptions = append(getPageDecryptions, *NewGetPageDecryptionResponse(&decryptions))
	}
	return &getPageDecryptions
}

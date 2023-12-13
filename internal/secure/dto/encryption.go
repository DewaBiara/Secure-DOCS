package dto

import (
	"github.com/DewaBiara/Secure-DOCS/pkg/entity"
)

type CreateEncryptionRequest struct {
	UserID     string `json:"userid" validate:"required"`
	Filename   string `json:"filename" validate:"required"`
	InputFile  string `json:"inputfile"`
	OutputFile string `json:"outputfile"`
}

func (u *CreateEncryptionRequest) ToEntity() *entity.Encryption {
	return &entity.Encryption{
		UserID:     u.UserID,
		Filename:   u.Filename,
		InputFile:  u.InputFile,
		OutputFile: u.OutputFile,
	}
}

type UpdateEncryptionRequest struct {
	ID       uint   `json:"id" validate:"required"`
	UserID   string `json:"userid" validate:"required"`
	Filename string `json:"filename" validate:"required"`
}

func (u *UpdateEncryptionRequest) ToEntity() *entity.Encryption {
	return &entity.Encryption{
		UserID:   u.UserID,
		Filename: u.Filename,
	}
}

type GetSingleEncryptionResponse struct {
	ID       uint   `json:"id"`
	UserID   string `json:"userid" validate:"required"`
	Filename string `json:"filename" validate:"required"`
}

func NewGetSingleEncryptionResponse(encryption *entity.Encryption) *GetSingleEncryptionResponse {
	return &GetSingleEncryptionResponse{
		ID:       encryption.ID,
		UserID:   encryption.UserID,
		Filename: encryption.Filename,
	}
}

type GetPageEncryptionResponse struct {
	ID       uint   `json:"id"`
	UserID   string `json:"userid" validate:"required"`
	Filename string `json:"filename" validate:"required"`
}

func NewGetPageEncryptionResponse(encryption *entity.Encryption) *GetPageEncryptionResponse {
	return &GetPageEncryptionResponse{
		ID:       encryption.ID,
		UserID:   encryption.UserID,
		Filename: encryption.Filename,
	}
}

type GetPageEncryptionsResponse []GetPageEncryptionResponse

func NewGetPageEncryptionsResponse(encryptions *entity.Encryptions) *GetPageEncryptionsResponse {
	var getPageEncryptions GetPageEncryptionsResponse
	for _, encryptions := range *encryptions {
		getPageEncryptions = append(getPageEncryptions, *NewGetPageEncryptionResponse(&encryptions))
	}
	return &getPageEncryptions
}

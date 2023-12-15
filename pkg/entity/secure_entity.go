package entity

import (
	"gorm.io/gorm"
)

type Encryption struct {
	gorm.Model
	UserID   string
	Filename string
}

type Encryptions []Encryption

type Key struct {
	gorm.Model
	PengirimID   string
	PenerimaID   string
	EncryptionID uint
	Key          string
}

type Keys []Key

type Decryption struct {
	gorm.Model
	UserID   string
	KeyID    uint
	Filename string
	Status   string
}

type Decryptions []Decryption

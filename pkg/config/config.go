package config

import (
	"os"

	"github.com/DewaBiara/INVM-System/pkg/entity"
	"github.com/google/uuid"
)

var (
	DefaultUser = &entity.User{
		ID:       uuid.New().String(),
		Username: "admin",
		Password: "admin",
		Role:     "admin",
	}
)

func LoadConfig() map[string]string {
	env := make(map[string]string)

	env["DB_HOST"] = os.Getenv("DB_HOST")
	env["DB_PORT"] = os.Getenv("DB_PORT")
	env["DB_USER"] = os.Getenv("DB_USER")
	env["DB_PASS"] = os.Getenv("DB_PASS")
	env["DB_NAME"] = os.Getenv("DB_NAME")
	env["PORT"] = os.Getenv("PORT")
	env["JWT_SECRET"] = os.Getenv("JWT_SECRET")

	return env
}

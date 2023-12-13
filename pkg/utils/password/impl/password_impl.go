package impl

import (
	"github.com/DewaBiara/Secure-DOCS/pkg/utils/password"
	"golang.org/x/crypto/bcrypt"
)

type PasswordFuncImpl struct {
}

func NewPasswordFuncImpl() password.PasswordFunc {
	return &PasswordFuncImpl{}
}

func (PasswordFuncImpl) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (PasswordFuncImpl) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

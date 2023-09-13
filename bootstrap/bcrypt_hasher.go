package bootstrap

import (
	"github.com/SawitProRecruitment/UserService/errors"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --name IBcryptHasher
type IBcryptHasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(hashedPassword, enteredPassword string) error
}
type bcryptPasswordHasher struct{}

func NewPasswordHasher() IBcryptHasher {
	return &bcryptPasswordHasher{}
}

func (b *bcryptPasswordHasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.ErrorGeneratePassword
	}
	return string(hashedPassword), nil
}

func (b *bcryptPasswordHasher) VerifyPassword(hashedPassword, enteredPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
}

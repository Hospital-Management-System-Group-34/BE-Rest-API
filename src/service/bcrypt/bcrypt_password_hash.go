package bcrypt

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
	"golang.org/x/crypto/bcrypt"
)

type bcryptPasswordHash struct {
}

func NewBcryptPasswordHash() application.PasswordHash {
	newBcryptPasswordHash := bcryptPasswordHash{}

	return &newBcryptPasswordHash
}

func (p *bcryptPasswordHash) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (p *bcryptPasswordHash) ComparePassword(plain string, encrypted string) error {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(plain))
}

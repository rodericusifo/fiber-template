package generator

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/rodericusifo/fiber-template/internal/pkg/config"
)

func GenerateHashFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.Env.PasswordHashingHashSalt)
	return string(bytes), err
}

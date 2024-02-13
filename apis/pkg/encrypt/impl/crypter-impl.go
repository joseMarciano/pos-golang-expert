package impl

import (
	"golang.org/x/crypto/bcrypt"
)

type DefaultCrypterImpl struct{}

func (crypter DefaultCrypterImpl) Encrypt(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	return string(bytes), err
}

func (crypter DefaultCrypterImpl) CompareValues(plaintext, ciphertext string) error {
	return bcrypt.CompareHashAndPassword([]byte(ciphertext), []byte(plaintext))
}

func (crypter DefaultCrypterImpl) Decrypt(ciphertext string) (string, error) {
	panic("Decrypt need to be implemented")
}

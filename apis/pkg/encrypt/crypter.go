package encrypt

import "github.com/joseMarciano/pos-golang-expert/apis/pkg/encrypt/impl"

type Crypter interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(ciphertext string) (string, error)
	CompareValues(plaintext, ciphertext string) error
}

func NewDefaultCrypter() Crypter {
	return &impl.DefaultCrypterImpl{}
}

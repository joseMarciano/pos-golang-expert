package user

import (
	"github.com/joseMarciano/pos-golang-expert/apis/pkg/encrypt"
	"github.com/joseMarciano/pos-golang-expert/apis/pkg/identifier"
)

type User struct {
	ID       identifier.ID `json:"id"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"-"` // password is ignored
}

func NewUser(name, email, password string) (*User, error) {
	pass, err := encrypt.NewDefaultCrypter().Encrypt(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       identifier.New(),
		Name:     name,
		Email:    email,
		Password: pass,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	return encrypt.NewDefaultCrypter().CompareValues(password, u.Password) == nil
}

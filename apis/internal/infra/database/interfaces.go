package database

import "github.com/joseMarciano/pos-golang-expert/apis/internal/entity/user"

type UserInterface interface {
	GetUser(id string) (user.User, error)
	FindByEmail(email string) (user.User, error)
}

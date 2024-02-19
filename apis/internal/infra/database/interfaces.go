package database

import (
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/product"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/user"
)

type UserInterface interface {
	GetUser(id string) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
	Create(user *user.User) error
}

type ProductInterface interface {
	Create(p *product.Product) error
	FindAll(page, limit int, sort string) ([]product.Product, error)
	FindByID(id string) (*product.Product, error)
	Update(p *product.Product) error
	Delete(id string) error
}

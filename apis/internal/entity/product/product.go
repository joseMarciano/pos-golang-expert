package product

import (
	"errors"
	"github.com/joseMarciano/pos-golang-expert/apis/pkg/identifier"
	"time"
)

type Product struct {
	ID        identifier.ID `json:"id"`
	Name      string        `json:"name"`
	Price     int           `json:"price"`
	CreatedAt time.Time     `json:"created_at"`
}

func NewProduct(name string, price int) *Product {
	return &Product{
		ID:        identifier.New(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errors.New("id is required")
	}

	if _, err := identifier.FromString(p.ID.String()); err != nil {
		return errors.New("invalid id")
	}

	if p.Name == "" {
		return errors.New("name is required")
	}

	if p.Price == 0 {
		return errors.New("price is required")
	}

	if p.CreatedAt.IsZero() {
		return errors.New("createdAt is required")
	}

	return nil
}

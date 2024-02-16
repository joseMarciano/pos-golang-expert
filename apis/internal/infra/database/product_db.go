package database

import (
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/product"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (p *ProductDB) Create(prod *product.Product) error {
	return p.DB.Create(prod).Error
}

func (p *ProductDB) FindAll(page, limit int, sort string) ([]product.Product, error) {
	var products []product.Product

	// Set offset and limit for pagination
	offset := (page - 1) * limit

	db := p.DB.Offset(offset).Limit(limit)

	// Add order clause if sort field provided
	if sort != "" {
		db = db.Order(sort)
	}

	// Get products from database
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductDB) FindByID(id string) (*product.Product, error) {
	pr := &product.Product{}
	err := p.DB.First(&pr, "id = ?", id).Error
	return pr, err
}

func (p *ProductDB) Update(prod *product.Product) error {
	_, err := p.FindByID(prod.ID.String())
	if err != nil {
		return err
	}

	return p.DB.Save(prod).Error
}

func (p *ProductDB) Delete(id string) error {
	_, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(&product.Product{}, "id = ?", id).Error
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{ID: uuid.New().String(), Name: name, Price: price}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	product := NewProduct("Mouse", 50.0)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Name = "Mouse Gamer"
	product.Price = 100.0

	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	prod, err := selectOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product found", prod)

	products, err := selectAllProducts(db)
	if err != nil {
		return
	}

	fmt.Println("Products found", products)

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product deleted")

	defer db.Close()
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	var product Product

	row := db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id)
	err := row.Scan(&product.ID, &product.Name, &product.Price) // passsing memory address to populate props
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	var products []Product

	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

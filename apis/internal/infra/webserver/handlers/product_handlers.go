package handlers

import (
	"encoding/json"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/dto"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/entity/product"
	"github.com/joseMarciano/pos-golang-expert/apis/internal/infra/database"
	"net/http"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

func (p *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productInput dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error on decode body"))
		return
	}

	pr := product.NewProduct(productInput.Name, productInput.Price)

	err = p.ProductDB.Create(pr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(pr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

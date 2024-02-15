package product

import "testing"

func TestNewProduct(t *testing.T) {

	name := "Test Product"
	price := 10

	p := NewProduct(name, price)

	if p.Name != name {
		t.Errorf("Name is not %s", name)
	}

	if p.Price != price {
		t.Errorf("Price is not %d", price)
	}

}

func TestValidate(t *testing.T) {
	p := Product{}
	err := p.Validate()

	if err == nil {
		t.Error("Validate should have returned error for empty product")
	}

}

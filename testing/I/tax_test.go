package tax

// go test .
// go test -v
import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectedTax := 5.0
	returnedTax := CalculateTax(amount)

	if returnedTax != expectedTax {
		t.Errorf("returnedTax %f want %f", returnedTax, 1.0)
	}
}

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

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expectedTax float64
	}

	table := []calcTax{
		{amount: 500.0, expectedTax: 5.0},
		{amount: 1000.0, expectedTax: 10.0},
		{amount: 1500.0, expectedTax: 10.0},
	}

	for _, tc := range table {
		returnedTax := CalculateTax(tc.amount)

		if returnedTax != tc.expectedTax {
			t.Errorf("returnedTax %f want %f", returnedTax, tc.expectedTax)
		}
	}
}

package tax

// go test .
// go test -v
// go test -coverprofile=coverage.out to see code coverage
// go test -html=coverage.out to see code coverage
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

// go test -bench=.
func BenchmarkCalculateTax(b *testing.B) {
	type calcTax struct {
		amount, expectedTax float64
	}

	table := []calcTax{
		{amount: 500.0, expectedTax: 5.0},
		{amount: 1000.0, expectedTax: 10.0},
		{amount: 1500.0, expectedTax: 10.0},
	}

	index := 0
	for i := 0; i < b.N; i++ {
		item := table[index]
		CalculateTax(item.amount)

		if index > 2 {
			index = 0
		}

		index++
	}
}

// go test -fuzz=.
func FuzzCalculateTax(f *testing.F) {
	testcases := []float64{
		-10.0,
		500.0,
		1000.0,
		1500.0,
		2000.0,
	}

	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("result %f want %f", result, 0.0)
		}

	})
}

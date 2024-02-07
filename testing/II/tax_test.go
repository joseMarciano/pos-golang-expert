package tax

import "testing"
import "github.com/stretchr/testify/assert"

func TestCalculateTax(t *testing.T) {
	returnedTax := CalculateTax(1000)
	assert.Equal(t, 10.0, returnedTax, "Tax is equal")
}

package tax

import (
	"errors"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestCalculateTax(t *testing.T) {
	returnedTax := CalculateTax(1000)
	assert.Equal(t, 10.0, returnedTax, "Tax is equal")
}

func TestCalculateTaxAndSave(t *testing.T) {
	repositoryMocked := &TaxRepositoryMock{}

	repositoryMocked.On("Save", 5.0).Return(nil)
	repositoryMocked.On("Save", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(10.0, repositoryMocked)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repositoryMocked)
	assert.NotNil(t, err)
}

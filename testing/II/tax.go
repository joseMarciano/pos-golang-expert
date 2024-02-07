package tax

type Repository interface {
	Save(tax float64) error
}

func CalculateTax(amount float64) float64 {
	if amount <= 0 {
		return 0
	}

	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0
	}

	if amount >= 1000 {
		return 10.0
	}
	return 5.0
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.Save(tax)
}

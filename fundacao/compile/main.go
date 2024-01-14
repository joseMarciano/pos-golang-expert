package main

import (
	UUID "github.com/google/uuid"
	"pos-golang-expert/utils"
)

// To compile go build main.go
// To compile to specific OS: GOOS=windows go build main.go
func main() {
	println(utils.Sum(1, 2))
	println(utils.PI)
	car := utils.Car{
		Brand: "Fiat",
		Model: "Uno",
		Year:  1990,
	}

	println(car.Brand)
	u := UUID.New()
	println(u.String())
}

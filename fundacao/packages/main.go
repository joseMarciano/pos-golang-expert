package main

import (
	UUID "github.com/google/uuid"
	"pos-golang-expert/utils"
)

func main() {
	println(utils.Sum(1, 2))
	println(utils.PI)
	car := utils.Car{
		Brand: "Fiat",
		Model: "Uno",
		Year:  1990,
	}

	//car.status = "VRUM VRUM" // private access, then I cant access here
	println(car.Brand)

	u := UUID.New()
	println(u)
}

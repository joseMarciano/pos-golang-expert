package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(10, 10))
	fmt.Println(sum(80, 79))
	fmt.Println(sumOrError(10, 10))
	fmt.Println(sumOrError(80, 79))
}

func sum(a, b int) (int, bool) {
	var result = a + b

	if result > 50 {
		return result, true
	}

	return result, false
}

func sumOrError(a, b int) (int, error) {
	var result = a + b

	if result > 50 {
		return result, errors.New("error on sum")
	}

	return result, nil
}

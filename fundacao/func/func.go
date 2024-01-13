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
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 1))
}

func sum(a, b int) (int, bool) {
	var result = a + b

	if result > 50 {
		return result, true
	}

	return result, false
}

func sumAll(values ...int) int {
	var total = 0
	for _, v := range values {
		total += v
	}

	return total
}

func sumOrError(a, b int) (int, error) {
	var result = a + b

	if result > 50 {
		return result, errors.New("error on sum")
	}

	return result, nil
}

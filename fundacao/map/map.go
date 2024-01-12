package main

import (
	"fmt"
	"strconv"
)

func main() {
	salary := map[string]float64{
		"John": 123.34,
		"Mike": 323.43,
	}

	for key, value := range salary {
		fmt.Printf("key: %s value: %f \r\n", key, value)
	}

	//We can use the make func to create a map
	var salaryMake = make(map[string]int)

	salaryMake["Travel"] = 123
	salaryMake["Franklin"] = 123
	fmt.Printf("Salary by make is %v", salaryMake)

	var arrayMap = make(map[string][]int)
	arrayMap["List_1"] = []int{1, 2, 3, 4, 5}
	arrayMap["List_2"] = []int{123, 4234, 234, 12, 23}

	for key, values := range arrayMap {
		var message string
		for _, v := range values {
			message += strconv.Itoa(v) + ", "
		}

		fmt.Printf("key: %s values: %s \r\n", key, message)
	}
}

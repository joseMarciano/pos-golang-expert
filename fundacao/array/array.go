package main

import "fmt"

func main() {
	var myArray [3]string // array has fixed size
	myArray[0] = "10"
	myArray[1] = "20"
	myArray[2] = "30"
	fmt.Printf("The array has %d of size \r\n", len(myArray))

	for i, v := range myArray {
		fmt.Printf("The value of index is %d and the value is %s \r\n", i, v)
	}

}

package main

import "fmt"

const A = "This is a constant"
const B = 1.5
const C = 99

func main() {
	fmt.Printf("The kind of A is %T and the value is %s \r\n", A, A)
	fmt.Printf("The kind of B is %T and the value is %f \r\n", B, B)
	fmt.Printf("The kind of C is %T and the value is %d \r\n", C, C)
}

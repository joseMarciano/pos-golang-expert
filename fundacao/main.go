package main

// Every variable declared outside a func is in a global scope and we can use in another packages

type ID int

var (
	name       string = "John"
	customType ID     = 1 // we can create a type
)

const constant = "Maria"

func main() {
	testing := "Man" // another way to declare variable
	println(name)
	println(customType)
	println(constant)
	println(testing)
}

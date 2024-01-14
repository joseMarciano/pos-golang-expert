package main

func main() {
	var a = 10
	var pointer *int = &a // here I am getting the memory address of "a"
	println(pointer)
	*pointer = 20 // here I am changing the value of "a"
	println(a)
	b := &a // here I am getting the memory address of "a". Here is shortcut for "var b *int = &a"
	println(b)
	*b = 30 // here I am changing the value of "a"
	println(a)
}

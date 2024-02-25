package main

func main() {
	/**
	This code will deadlock because will be waiting the channel be empty to put "World"
	ch := make(chan string)
	ch <- "Hello"
	ch <- "World"
	print(<-ch)
	print(<-ch)
	*/

	ch := make(chan string, 2) // this code will work just fine because I have a buffered channel
	ch <- "Hello "
	ch <- "World"
	print(<-ch)
	print(<-ch)
}

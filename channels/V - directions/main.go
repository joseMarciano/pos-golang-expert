package main

import "fmt"

func main() {
	var ch = make(chan string)
	go receive("A", ch)
	read(ch)

}

func receive(name string, ch chan<- /*the <- means that this func can just publish in the channel*/ string) {
	ch <- name
}

func read(ch <-chan /*the <- means that this func can just get data from the channel*/ string) {
	fmt.Println(<-ch)
}

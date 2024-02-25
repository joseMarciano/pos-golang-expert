package main

import "fmt"

func main() {
	var channel = make(chan int)
	go publish(channel)

	receiver(channel)
}

func receiver(ch chan int) {
	for x := range ch {
		fmt.Printf("%d, ", x)
	}
}

func publish(ch chan int) {
	defer close(ch) // close the channel to avoid deadlock. When this func finish, if I do not close, the receiver func will be waiting forever
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

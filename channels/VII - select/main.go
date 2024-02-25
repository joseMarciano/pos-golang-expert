package main

import "time"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	for i := 0; i < 3; i++ {
		select {
		case <-ch1: // case msg := <- ch1
			println("ch1")
		case <-ch2:
			println("ch2")
		case <-time.After(3 * time.Second):
			println("timeout")
			//default:
			//	println("default")
		}
	}
}

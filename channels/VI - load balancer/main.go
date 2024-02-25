package main

import (
	"fmt"
	"time"
)

func main() {
	var data = make(chan int)
	defer close(data)
	var workerQnt = 100

	for i := 0; i < workerQnt; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}

}

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(1 * time.Second)
	}
}

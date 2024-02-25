package main

import (
	"fmt"
	"sync"
)

const INTERACTIONS_TIMES = 10

// In this example, the publish and receiver are running in different goroutines. The waiting group is necessary in this case to not stop process in the main thread. So, the main thread will wait until all data are transfered to receiver
func main() {
	var channel = make(chan int)
	var wg = sync.WaitGroup{}
	wg.Add(INTERACTIONS_TIMES)

	go publish(channel)
	go receiver(channel, &wg)

	wg.Wait()
}

func receiver(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("%d, ", x)
		wg.Done()
	}
}

func publish(ch chan int) {
	defer close(ch) // close the channel to avoid deadlock. When this func finish, if I do not close, the receiver func will be waiting forever
	for i := 0; i < INTERACTIONS_TIMES; i++ {
		ch <- i
	}
}

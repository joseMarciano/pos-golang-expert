package main

import (
	"fmt"
	"sync"
	"time"
)

const NUMBER_OF_OPERATIONS = 15
const NUMBER_OF_FOR_LOOP = 5

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(NUMBER_OF_OPERATIONS)

	//thread1
	go task("A", &waitGroup)
	//thread2
	go task("B", &waitGroup)
	//thread3
	go task("C", &waitGroup)

	waitGroup.Wait() // will wait until arrived 0
	print("Done")
}

func task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < NUMBER_OF_FOR_LOOP; i++ {
		fmt.Printf("Running task in thread %s\n", name)
		time.Sleep(1 * time.Second)
		waitGroup.Done() // will reduce the number every done task
	}
}

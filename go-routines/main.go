package main

import (
	"fmt"
	"time"
)

func main() {
	//thread1
	go task("A")
	//thread2
	go task("B")
	//thread3
	go task("C")

	time.Sleep(20 * time.Second)
}

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Running task in thread %s\n", name)
		time.Sleep(10 * time.Second)
	}
}

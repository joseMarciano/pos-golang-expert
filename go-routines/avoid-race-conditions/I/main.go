package main

// if we want to know possible race conditions, we need to run go run --race main.go(just in developer environment)

import (
	"sync"
	"sync/atomic"
	"time"
)

var mutexCounter uint64 = 0
var mtx = sync.Mutex{}

var atomicCounter uint64 = 0

func main() {
	for i := 0; i < 10000; i++ {
		go avoidRaceConditionsUsingMutex()
		go avoidRaceConditionsUsingAtomicPackage()
	}

	time.Sleep(15 * time.Second)
	println(mutexCounter)
	println(atomicCounter)
}

func avoidRaceConditionsUsingAtomicPackage() {
	atomic.AddUint64(&atomicCounter, 1)
}

func avoidRaceConditionsUsingMutex() {
	mtx.Lock()
	mutexCounter++
	mtx.Unlock()
}

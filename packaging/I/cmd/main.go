package main

import (
	"github.com/devfullcycle/goexpert/packaging/1/math"
	"log"
)

func main() {
	m := math.Math{
		A: 10,
		B: 20,
	}
	log.Printf("Math lib was imported succesfully:  %v", m.Add())
}

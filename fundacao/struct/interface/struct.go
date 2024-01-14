package main

import "fmt"

type Person interface {
	Inactivate() // every struct that have Inactivate method automatically implements Person
}

type Client struct {
	Name   string
	Age    uint8
	Active bool
}

func (c Client) Inactivate() {
	c.Active = false
}

func printHere(p Person) {
	fmt.Println(p)
}

func main() {

	client := Client{
		Name:   "John",
		Age:    12,
		Active: true,
	}

	printHere(client) // client is a Person
}

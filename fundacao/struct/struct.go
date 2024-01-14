package main

import "fmt"

type Client struct {
	Name   string
	Age    uint8
	Active bool
}

func main() {

	client := Client{
		Name:   "John",
		Age:    10,
		Active: false,
	}

	client.Active = true

	fmt.Println(client)
}

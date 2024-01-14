package main

import "fmt"

type Address struct {
	Street string
	Number uint
}

type Client struct {
	Name    string
	Age     uint8
	Active  bool
	Address // We can use Adresss like that. So the client will be like client.Street
	//Address Address // We can use Adresss like that. So the client will be like client.Address.Street
}

func main() {

	client := Client{
		Name:   "John",
		Age:    12,
		Active: false,
		Address: Address{
			Street: "Frei",
			Number: 213,
		},
	}

	client.Active = true

	fmt.Println(client)
}

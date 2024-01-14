package main

import "fmt"

type Client struct {
	Name   string
	Age    uint8
	Active bool
}

func (c *Client) inactivate() *Client {
	c.Active = false
	return c
}

func (c Client) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Active: %t", c.Name, c.Age, c.Active)
}

func main() {

	client := Client{
		Name:   "John",
		Age:    12,
		Active: true,
	}

	client.inactivate()

	fmt.Println(client)
}

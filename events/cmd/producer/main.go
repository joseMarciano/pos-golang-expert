package main

import (
	"github.com/devfullcycle/fcutils/pkg/rabbitmq"
	"strconv"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 10; i++ {
		rabbitmq.Publish(ch, "Testando "+strconv.Itoa(i))
	}
}

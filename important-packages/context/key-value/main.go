package main

import (
	"context"
	"fmt"
)

type RandomData struct {
	token     string
	bookingId int
}

func main() {
	randomData := RandomData{
		token:     "some-value",
		bookingId: 123,
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "ctx-value", randomData)
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	randomData := ctx.Value("ctx-value")
	fmt.Println(randomData)
	// do something with token and bookingId
}

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	background := context.Background()
	ctx, cancel := context.WithTimeout(background, time.Second*3)
	defer cancel() // good practice  to call cancel when you're done with the context
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	//select is a
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked")
		return
	}

}

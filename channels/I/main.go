package main

// thread1
func main() {

	channel := make(chan string) // empty channel

	//thread2
	go func() {
		channel <- "Hello"   // put value in channel
		channel <- "Hello 2" // this value will be put in the channel just when the "Hello" value be consumed
	}()

	var message = <-channel // wait until the channel is filled and put in variable
	println(message)
	message = <-channel
	println(message)

}

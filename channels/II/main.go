package main

//
//func main() {
//	var forever = make(chan bool)
//
//	<-forever // this will a dead lock because the channel is always empty and will be wait forever
//}

func main() {
	var forever = make(chan bool)

	go func() {
		for i := 0; i < 50; i++ {
			println(i)
		}
		forever <- true
	}()

	<-forever // in this case, this thread will be waiting until de anonymous func finish and set true to the channel
}

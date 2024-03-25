package main

func myPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("recovered from ", r)
		}
	}()
	myPanic()
}

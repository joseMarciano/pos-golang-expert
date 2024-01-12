package main

import "fmt"

func main() {

	mySlice := []int{2, 3, 4, 5, 6, 7, 9}
	fmt.Printf("len=%d  cap=%d var=%v \n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("len=%d  cap=%d var=%v \n", len(mySlice[:0]), cap(mySlice[:0]), mySlice[:0])
	fmt.Printf("len=%d  cap=%d var=%v \n", len(mySlice[:1]), cap(mySlice[:1]), mySlice[:1])
	fmt.Printf("len=%d  cap=%d var=%v \n", len(mySlice[3:]), cap(mySlice[3:]), mySlice[3:])

	mySlice = append(mySlice, 123)
	fmt.Printf("\r\nlen=%d  cap=%d var=%v \n", len(mySlice[3:]), cap(mySlice[3:]), mySlice[3:])
	fmt.Printf("len=%d  cap=%d var=%v \n", len(mySlice[:3]), cap(mySlice[:3]), mySlice[:3])
}

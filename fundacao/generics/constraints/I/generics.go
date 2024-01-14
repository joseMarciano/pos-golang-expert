package main

type MyNumber int

type Number interface { // I can use constraints like Number
	int | float64
}

//type Number interface { To MyNumber type be acceptable in Number type, we need put ~ in front of int
//	~int | float64
//}

func main() {
	m := map[string]int{"Wesley": 1000}
	//m1 := map[string]MyNumber{"Wesley": 1000}
	println(sum(m))
	//println(sum(m1)) // for this code run, web need use the commented type Number, because m1 is type of MyNumber
}

func sum[T Number](mp map[string]T) T {
	var sum T
	for _, value := range mp {
		sum += value
	}
	return sum
}

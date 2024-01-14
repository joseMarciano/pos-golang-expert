package main

type Number interface { // I can use constraints like Number
	int | float64
}

type AnyContract interface {
}

func main() {
	m := map[string]int{"Wesley": 1000}
	//m1 := map[string]MyNumber{"Wesley": 1000}
	println(sum(m, "Testing"))
	println(sum(m, 0))
}

func sum[T Number, B AnyContract](mp map[string]T, b B) (T, B) {
	var sum T
	for _, value := range mp {
		sum += value
	}
	return sum, b
}

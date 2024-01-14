package main

func main() {
	m := map[string]int{"Wesley": 1000}
	println(sum(m))
}

func sum[T int | float64](mp map[string]T) T {
	var sum T
	for _, value := range mp {
		sum += value
	}
	return sum
}

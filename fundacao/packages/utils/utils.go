package utils

type Car struct {
	Brand  string
	Model  string
	Year   int
	status string
}

func (c *Car) Accelerate() string {
	c.status = "VRUM VRUM"
	return c.status
}

var PI = 3.141516 // public access because init with upperCase

func Sum[T int | float64](a, b T) T { // public access because init with upperCase
	return a + b
}

func sub[T int | float64](a, b T) T { // private access because init with lower case
	return a - b
}

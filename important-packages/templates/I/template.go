package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name  string
	Price float64
}

func main() {
	course := Course{
		Name:  "Go",
		Price: 45.0,
	}

	tmp := template.New("Template")
	tmp, _ = tmp.Parse("Name: {{.Name}} Price: {{.Price}}")
	tmp.Execute(os.Stdout, course)
}

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

	t := template.Must(template.New("Template").Parse("Name: {{.Name}} Price: {{.Price}}"))
	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}

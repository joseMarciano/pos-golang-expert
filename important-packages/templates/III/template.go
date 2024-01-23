package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name  string
	Price float64
}

type Courses []Course

func main() {

	t := template.Must(template.New("main.html").ParseFiles("/home/usuario/JOSEPAULO/pos-golang-expert/important-packages/templates/IV/main.html"))
	err := t.Execute(os.Stdout, Courses{
		{"Go", 5},
		{"Java", 10},
		{"Python", 15},
		{"JavaScript", 20},
		{"TypeScript", 25},
		{"PHP", 30},
		{"Ruby", 35},
	})
	if err != nil {
		panic(err)
	}
}

package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 100
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {

	p := person{
		Name: "Rethabile Mupereri",
		Age:  3,
	}

	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
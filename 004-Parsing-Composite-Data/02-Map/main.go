package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {

	planetDistances := map[string]float64{
		"Mercury": 0.39,
		"Venus": 0.72,
		"Earth": 1,
		"Mars": 1.52,
		"Jupiter": 5.2,
		"Saturn": 9.54,
		"Uranus": 19.2,
		"Neptune": 30.06,
	}


	err := tpl.Execute(os.Stdout, planetDistances)
	if err != nil {
		log.Fatalln(err)
	}

}
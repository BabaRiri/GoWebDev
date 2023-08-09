package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

//template.Must handles error for us
func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	//We parsed "BabaRiri" to the placeholder on our html file
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", "BabaRiri")
	if err != nil {
		log.Fatalln(err)
	}
}
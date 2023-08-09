package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

//tpl is package level meaning can be accessed at package level
var tpl *template.Template

//This function makes sure that all my files are parsed to my template before the main function runs
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	//executing the template to the standard output(terminal)
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//creating a new file
	newFile, err := os.Create("index.html")
	if err != nil {
	 log.Fatalln(err)
	}
	defer newFile.Close()

	//executing the template to the new file
	err = tpl.Execute(newFile, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//Parsing files to the template
	/*
	tpl, err = tpl.ParseFiles("one.gohtml", "two.gohtml", "three.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	*/

	//executing each file to the standard output
	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n")

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n")

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n")

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.html", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
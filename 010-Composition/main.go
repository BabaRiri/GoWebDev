package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"BUSN461", "STRATEGIC PLANNNING AND MANAGEMENT", "4"},
				{"COMP 415 ", "ARTIFICIAL INTELLIGENCE", "4"},
				{"EE 419", "WIRELESS COMMUNICATION", "4"},
				{"ECE 408", "DIGITAL SYSTEM PROCESSING", "4"},
				{"ECON413", "ENGINEERING ECONOMICS", "4"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"COMP 464", "INTERNET PROGRAMMING", "5"},
				{"EE431", "PRINCIPLES OF DIGITAL IMAGE PROCESSING", "5"},
				{"COMP448", "ARTIFICIAL NEURAL NETWORKS", "5"},
				{"ENG434", "ENGINEERING ETHICS", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
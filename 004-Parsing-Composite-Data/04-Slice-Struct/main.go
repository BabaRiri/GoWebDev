package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Company struct {
	Name    string
	Slogan  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {

	apple := Company{
		Name:   "Apple",
		Slogan: "Think Different",
	}

	amazon := Company{
		Name:   "Amazon",
		Slogan: "Work Hard, Have Fun, Make History",
	}

	facebook := Company{
		Name:   "Facebook",
		Slogan: "Move Fast and Break Things",
	}

	netflix := Company{
		Name:   "Netflix",
		Slogan: "See What's Next",
	}

	google := Company{
		Name:   "Google",
		Slogan: "Don't Be Evil",
	}

	companies := []Company{apple, amazon, facebook, netflix, google}

	err := tpl.Execute(os.Stdout, companies)
	if err != nil {
		log.Fatalln(err)
	}
}

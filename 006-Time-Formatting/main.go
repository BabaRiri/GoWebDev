package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.html"))
}

func monthDayYear(t time.Time) string {
	return t.Format("2006-01-02")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
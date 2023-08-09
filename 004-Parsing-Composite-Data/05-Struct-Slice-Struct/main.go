package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Item struct {
	Name  string
	Price float64
}

type Category struct {
	Name  string
	Items []Item
}

type Store struct {
	Name      string
	Categories []Category
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {

	itemsElectronics := []Item{
		{Name: "Laptop", Price: 999.99},
		{Name: "Smartphone", Price: 599.99},
		{Name: "Headphones", Price: 149.99},
	}

	itemsClothing := []Item{
		{Name: "T-Shirt", Price: 29.99},
		{Name: "Jeans", Price: 49.99},
		{Name: "Shoes", Price: 79.99},
	}

	categories := []Category{
		{Name: "Electronics", Items: itemsElectronics},
		{Name: "Clothing", Items: itemsClothing},
	}

	store := Store{
		Name:      "SuperMart",
		Categories: categories,
	}

	err := tpl.Execute(os.Stdout, store)
	if err != nil {
		log.Fatalln(err)
	}
}

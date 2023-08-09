package main

import (
	"log"
	"os"
	"text/template"
)

type Demographics struct {
    Population int
    GrowthRate float64
    BirthRate float64
    DeathRate float64
    LifeExpectancy float64
    FertilityRate float64
    InfantMortalityRate float64
    NetMigrationRate float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {

	zimbabweDemographics := Demographics{
		Population: 15993524,
		GrowthRate: 1.95,
		BirthRate: 33.07,
		DeathRate: 8.76,
		LifeExpectancy: 63.32,
		FertilityRate: 3.89,
		InfantMortalityRate: 28.53,
		NetMigrationRate: -4.83,
	}


	err := tpl.Execute(os.Stdout, zimbabweDemographics)
	if err != nil {
		log.Fatalln(err)
	}

}
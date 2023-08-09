package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type programmer struct {
	person
	code bool
}

func (p person) pSpeak() {
	fmt.Printf("My name is %s\n", p.firstName)
}

func (pg programmer) pgSpeak() {
	fmt.Printf("My name is %s, and i can code in GO!\n", pg.firstName)
}

func main() {
	p1 := person {
		firstName: "Nthabiseng",
		lastName: "Mupereri",
	}

	p2 := programmer {
		person{
			firstName: "Seward",
			lastName: "Mupereri",
		},
		true,
	}

	fmt.Printf("My last name is %s\n", p1.lastName)
	p1.pSpeak()

	fmt.Printf("My last name is %s\n", p2.lastName)
	p2.pgSpeak()
	p2.pSpeak()

}

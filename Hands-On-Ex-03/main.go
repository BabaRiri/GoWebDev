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

func (p person) Speak() string {
	return fmt.Sprintf("My name is %s\n", p.firstName)
}

func (pg programmer) Speak() string {
	return fmt.Sprintf("My name is %s, and i can code in GO!\n", pg.firstName)
}

type human interface {
	Speak() string
}
func talk(h human) {
	switch x := h.(type) {
		case person:
			fmt.Println(x.firstName)
		case programmer:
			fmt.Println(x.firstName)
		default:
			fmt.Println("Who are you?")
	}
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
	p1.Speak()

	fmt.Printf("My last name is %s\n", p2.lastName)
	p2.Speak()

	talk(p1)
	talk(p2)

}

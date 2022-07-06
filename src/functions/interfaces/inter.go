package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	fname string
	lname string
}

type agent struct {
	person
	license_to_kill bool
}

func (a agent) speak() {
	fmt.Println("Agent:", a.fname, a.lname, a.license_to_kill)
}

func (p person) speak() {
	fmt.Println("Person:", p.fname, p.lname)
}

func multipleArgTypes(h human) {
	h.speak()
}

func main() {
	james_bond := agent{
		person: person{
			fname: "James",
			lname: "Bond",
		},
		license_to_kill: true,
	}

	joaquim := person{
		fname: "Joaquim",
		lname: "Teixeira",
	}

	defer fmt.Println("Done!")

	multipleArgTypes(james_bond)
	multipleArgTypes(joaquim)
}

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int8
}

type agent struct {
	person
	license_to_kill bool
}

func main() {
	james_bond := agent{
		person: person{
			name: "James Bond",
			age:  35,
		},
		license_to_kill: true,
	}

	fmt.Println(james_bond.name, james_bond.person.age, james_bond.age, james_bond.license_to_kill)
}

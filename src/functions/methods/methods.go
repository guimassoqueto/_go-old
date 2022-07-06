package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type agent struct {
	person
	licence_to_kill bool
}

func (a agent) speak() {
	if a.licence_to_kill {
		defer fmt.Println("I have license to kill.")
	} else {
		defer fmt.Println("I do not have license to kill.")
	}
	fmt.Println("I am", a.name)
}

func main() {
	james_bond := agent{
		person: person{
			name: "James Bond",
			age:  35,
		},
		licence_to_kill: true,
	}

	mega_gaymen := agent{
		person: person{
			name: "Gaymen King",
			age:  35,
		},
		licence_to_kill: false,
	}

	james_bond.speak()
	mega_gaymen.speak()
}

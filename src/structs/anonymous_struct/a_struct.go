package main

import (
	"fmt"
)

func main() {
	person := struct {
		name string
		age  int
	}{
		"Joaquim",
		55,
	}

	fmt.Println(person.name, person.age)
}

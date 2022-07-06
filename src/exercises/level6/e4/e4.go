package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi, my name is", p.first, p.last, "and i am ", p.age, "years old.")
}

func main() {
	jordi_polla := person{
		first: "Jordi",
		last:  "El Niño Polla",
		age:   27,
	}

	jordi_polla.speak()
}

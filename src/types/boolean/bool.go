package main

import "fmt"

var my_boolean1 bool

func main() {
	fmt.Printf("1. %t\n", my_boolean1)
	my_boolean1 = true
	fmt.Printf("2. %t\n", my_boolean1)

	a := 55
	b := 27
	c := a == b

	fmt.Printf("3. %t\n", c)
}

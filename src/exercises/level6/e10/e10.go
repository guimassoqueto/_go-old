package main

import (
	"fmt"
)

func out() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	a := out()
	a()
	a()
	a()

	b := out()
	b()

	fmt.Println(a())
	fmt.Println(b())
	
}

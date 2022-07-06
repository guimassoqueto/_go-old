package main

import "fmt"

func outer() func() string {
	return func() string {
		return "Inner Func"
	}
}

func main() {
	a := outer()
	b := a()

	fmt.Println(b)
}

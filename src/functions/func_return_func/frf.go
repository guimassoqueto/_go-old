package main

import "fmt"

func main() {
	x := func() func() int {
		return func() int {
			return 32
		}
	}

	fmt.Println(x()())
}

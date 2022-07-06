package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Defer!")
	}()

	test := func() {
		fmt.Printf("function called\n")
	}

	test()
}

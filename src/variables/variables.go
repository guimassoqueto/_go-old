package main

import "fmt"

var out int = 255

func main() {
	fmt.Print(out)

	out := 250
	for i := out; i <= 256; i++ {
		fmt.Println(i)
	}
	fn()
}

func fn() {
	fmt.Println("fn: ", out)
}

package main

import "fmt"

var s string

func main() {
	s = "Hello"
	b := []byte(s)

	fmt.Printf("%s, %b\n", s, b)
}

package main

import "fmt"

const (
	a int     = 54
	b float64 = 2.3658
	c string  = "Test"
)

const d bool = true

func main() {
	b := a
	fmt.Printf("%d\n", b+2)
}

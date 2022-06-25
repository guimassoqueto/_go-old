package main

import "fmt"

/* assigned consts */
const (
	a int    = 42
	b string = "hello"
	c bool   = false
)

/* unssigned consts */
const (
	d = 42
	e = "hello"
	f = false
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

package main

import (
	"fmt"
	"math/rand"
)

func foo(max int) int {
	return rand.Intn(max)
}

func bar(s string) (int, string) {
	return len(s), s
}

func main() {
	_int := foo(10)
	in, str := bar("guilherme")

	fmt.Println(_int)
	fmt.Println(in, str)
}

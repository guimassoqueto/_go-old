package main

import "fmt"

func variadic(vals ...int) []int {
	return vals
}

func main() {
	test := variadic(1, 3, 5, 7, 9)

	for _, val := range test {
		fmt.Println(val)
	}
}

package main

import (
	"fmt"
)

func foo(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

func bar(vals []int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	numbers := []int{10, 20, 25, 35}

	result_foo := foo(numbers...)
	fmt.Println(result_foo)

	result_bar := bar(numbers)
	fmt.Println(result_bar)
}

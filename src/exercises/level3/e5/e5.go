package main

import (
	"fmt"
)

const divisor int = 4

func main() {
	for i := 10; i <= 100; i++ {
		var rest int = i % divisor
		fmt.Println(rest)
	}
}

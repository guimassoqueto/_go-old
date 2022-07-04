package main

import (
	"fmt"
)

var number int8 = 1

func main() {
	switch number {
	case 4:
		fmt.Printf("%d is four\n", number)
	case 0:
		fmt.Printf("%d is zero\n", number)
	default:
		fmt.Printf("%d is neither four nor zero\n", number)
	}
}

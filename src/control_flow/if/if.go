package main

import (
	"fmt"
)

var number int8 = 0

func main() {
	if number < 0 {
		fmt.Printf("%d is less than 0\n", number)
	} else if number == 0 {
		fmt.Printf("%d is equal 0\n", number)
	} else {
		fmt.Printf("%d is greater than 0\n", number)
	}
}

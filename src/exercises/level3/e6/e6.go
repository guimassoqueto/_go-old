package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	fmt.Print("Type a number: ")
	fmt.Scanln(&num)

	number, err := strconv.Atoi(num)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if number%2 == 0 {
		fmt.Printf("%d is even\n", number)
	} else {
		fmt.Printf("%d is odd\n", number)
	}
}

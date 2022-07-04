package main

import (
	"fmt"
)

func main() {
	switch {
	case -1 < 0:
		fmt.Println("-1 less than 0")
	case 5 == 4:
		fmt.Println("5 is equal 4")
	default:
		fmt.Println("default")
	}
}

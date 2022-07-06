package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Done!")

	func(number int) {
		fmt.Println(number)
	}(42)
}

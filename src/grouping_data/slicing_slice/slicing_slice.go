package main

import (
	"fmt"
)

func main() {
	var slice = []string{"a", "b", "c", "d", "e"}

	fmt.Println(slice[1:4])

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "b", "c"}
	slice = append(slice, "d")

	for index, value := range slice {
		fmt.Printf("%d - %s\n", index, value)
	}
}

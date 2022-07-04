package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 10, 15)

	slice = append(slice, 11, 12, 13, 14, 15, 17)
	fmt.Println(slice)
}

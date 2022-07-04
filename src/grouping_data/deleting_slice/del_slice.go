package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice = append(slice[3:6], slice[8:]...)

	fmt.Println(slice)
}

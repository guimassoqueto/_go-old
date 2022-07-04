package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}

	slice = append(slice, 4, 5, 6, 7)

	sliceOne := []string{"a", "b", "c"}
	sliceTwo := []string{"x", "y", "z"}
	sliceThree := []string{"1", "2", "3"}

	sliceFour := append(sliceOne[:1], sliceTwo...)
	fmt.Println(sliceFour)
	sliceFour = append(sliceFour, sliceThree...)
	fmt.Println(sliceFour)
}

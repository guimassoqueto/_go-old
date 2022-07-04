package main

import "fmt"

func main() {
	slice := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice = append(slice[1:4], slice[7:]...)

	fmt.Println(slice)
}

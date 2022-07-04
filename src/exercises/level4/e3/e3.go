package main

import "fmt"

func main() {
	slice := [15]int{40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54}

	fmt.Println(slice[2:7])
	fmt.Println(slice[7:12])
	fmt.Println(slice[4:9])
	fmt.Println(slice[3:8])
}

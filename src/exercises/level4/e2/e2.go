package main

import "fmt"

func main() {
	slice := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

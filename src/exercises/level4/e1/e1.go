package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 100
	arr[1] = 101
	arr[2] = 102
	arr[3] = 103
	arr[4] = 104

	for _, v := range arr {
		fmt.Println(v)
	}
}

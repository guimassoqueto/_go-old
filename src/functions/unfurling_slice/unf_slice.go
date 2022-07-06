package main

import "fmt"

func displayNumbers(numbers ...int) {
	for _, num := range numbers {
		fmt.Println(num)
	}
}

func main() {
	fnum := []int{1, 2}
	displayNumbers(fnum...)
}

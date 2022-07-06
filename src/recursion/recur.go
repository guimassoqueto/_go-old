package main

import "fmt"

func factorial(numb uint) uint {
	if numb == 0 || numb == 1 {
		return 1
	}

	return numb * factorial(numb-1)
}

func main() {
	val := factorial(20)

	fmt.Println(val)
}

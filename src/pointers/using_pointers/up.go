package main

import (
	"fmt"
)

var x uint8 = 42

func main() {
	fmt.Println("x before changeByVal", x)
	changeByVal(x)
	fmt.Printf("x after changeByVal %d\n", x)

	fmt.Printf("\nx before changeByRef %d\n", x)
	changeByRef(&x)
	fmt.Println("x after changeByRef", x)
}

func changeByVal(y uint8) {
	y = 66
	fmt.Println("x inside changeByVal", y)
}

func changeByRef(y *uint8) {
	*y = 66
	fmt.Println("x inside changeByRef", *y)
}

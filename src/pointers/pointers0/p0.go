package main

import (
	"fmt"
)

var a uint8 = 255

func main() {
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println("\nb's location:", b)
	fmt.Println("b's value:", *b)
	*b = 100

	fmt.Println("\na's value: ", a)
	fmt.Printf("a's type: %T\n", a)
	fmt.Printf("a's reference type: %T\n", &a)
	fmt.Println("b's location:", b)
	fmt.Println("b's value:", *b)
}

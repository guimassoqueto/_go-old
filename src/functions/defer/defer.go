package main

import "fmt"

func foo() {
	fmt.Println("Foo...")
}

func bar() {
	fmt.Println("Bar...")
}

func test() {
	fmt.Println("Test...")
}

func main() {
	defer bar()
	defer foo()
	test()
	test()
	test()
}

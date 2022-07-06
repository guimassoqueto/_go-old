package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous func callled	without args")
	}()

	func(i int) {
		fmt.Println("Answer to meaning of life:", i)
	}(42)
}

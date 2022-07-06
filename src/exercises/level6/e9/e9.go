package main

import "fmt"

func test(number int, callback func(number_plus_ten int)) {
	callback(number + 10)
}

func main() {
	test(10, func(num int) {
		fmt.Println(num)
	})
}

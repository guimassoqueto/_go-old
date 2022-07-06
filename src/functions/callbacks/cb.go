package main

import "fmt"

func sum_evens(max int, callback func(nums int)) {
	even_sum := 0

	for i := 0; i <= max; i++ {
		if i%2 != 0 {
			even_sum += i
		}
	}
	callback(even_sum)
}

func main() {
	sum_evens(11, func(total int) {
		fmt.Println("total: ", total)
	})
}

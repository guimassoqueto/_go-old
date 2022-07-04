package main

import (
	"fmt"
)

func main() {
	var birth_year int16 = 1992

	for birth_year <= 2022 {
		fmt.Println(birth_year)
		birth_year++
	}
}

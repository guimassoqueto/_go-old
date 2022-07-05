package main

import (
	"fmt"
)

func main() {
	ferrari := struct {
		color string
		model [2]string
		price map[string]int
	}{
		color: "red",
		model: [2]string{"F50", "F50_1998"},
		price: map[string]int{
			"minimum": 1_000_000,
			"maximum": 1_500_000,
		},
	}

	fmt.Println(ferrari.color, ferrari.model, ferrari.price)
}

package main

import (
	"fmt"
)

type person struct {
	fname       string
	lname       string
	favIceCream []string
}

func main() {
	joaquim := person{
		fname: "Joaquim",
		lname: "Teixeira",
		favIceCream: []string{
			"Chocolate",
			"Vanilla",
		},
	}

	fmt.Println(joaquim.fname, joaquim.lname, joaquim.favIceCream)
}

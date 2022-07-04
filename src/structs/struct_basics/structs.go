package main

import "fmt"

type person struct {
	fname string
	lname string
	age   uint8
}

func main() {
	joaquim := person{
		fname: "Joaquim",
		lname: "Teixeira",
	}
	joaquim.age = 65

	fmt.Println(joaquim.age, joaquim.fname, joaquim.lname)
}

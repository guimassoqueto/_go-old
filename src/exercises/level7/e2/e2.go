package main

import "fmt"

type agent struct {
	fname string
	lname string
}

func main() {
	bond := agent{
		fname: "James",
		lname: "Bond",
	}
	fmt.Println("before:", bond)

	changeByRef(&bond, "Gaymen", "Retarded")

	fmt.Println("after", bond)
}

func changeByRef(a *agent, new_fname string, new_lname string) {
	(*a).fname = new_fname
	(*a).lname = new_lname
}

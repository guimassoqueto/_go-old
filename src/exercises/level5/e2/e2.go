package main

import "fmt"

type person struct {
	fname       string
	lname       string
	favIceCream []string
}

func main() {
	p1 := person{
		fname:       "José",
		lname:       "Silva",
		favIceCream: []string{"Choco", "Vanilla", "Bacon"},
	}

	p2 := person{
		fname:       "Maria",
		lname:       "Souza",
		favIceCream: []string{"Rum", "Vodka", "Whisky"},
	}

	people := map[string]person{}

	people[p1.fname] = p1
	people[p2.fname] = p2

	for people_key, people_value := range people {
		fmt.Println(people_key)
		fmt.Printf("\t%v\n", people_value.fname)
		fmt.Printf("\t%v\n", people_value.lname)

		for _, ice_cream := range people_value.favIceCream {
			fmt.Printf("\t\t%v\n", ice_cream)
		}
		fmt.Printf("\n")
	}
}

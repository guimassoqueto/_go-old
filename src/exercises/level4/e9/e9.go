package main

import "fmt"

func main() {
	person := map[string][]string{}

	person["bond_james"] = []string{"Martini", "Women", "Kill Peope"}
	person["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	person["no_dr"] = []string{"Being Evil", "Ice Cream", "Sunsets"}

	// checks if the key exists
	if _, ok := person["no_dr"]; ok {
		delete(person, "no_dr")
	}

	for name, fav_things_arr := range person {
		fmt.Println(name, fav_things_arr)
	}
}

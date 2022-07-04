package main

import "fmt"

func main() {
	person := map[string][]string{}

	person["bond_james"] = []string{"Martini", "Women", "Kill Peope"}
	person["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	person["no_dr"] = []string{"Being Evil", "Ice Cream", "Sunsets"}

	for name, fav_things_arr := range person {
		fmt.Println(name)
		for fav_thing_index, fav_thing := range fav_things_arr {
			fmt.Printf("\t %v %v\n", fav_thing_index, fav_thing)
		}
		fmt.Println("")
	}
}

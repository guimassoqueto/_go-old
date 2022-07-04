package main

import "fmt"

func main() {
	map_ := map[string]int{}

	map_["lula"] = 22
	map_["bolsonaro"] = 13

	fmt.Println(map_)
	fmt.Println(map_["lula"])

	delete(map_, "bolsonaro")

	for i, v := range map_ {
		fmt.Printf("%v ---> %v\n", i, v)
	}
}

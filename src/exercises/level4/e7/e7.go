package main

import "fmt"

func main() {
	person := map[string][2]string{
		"Jorge": {"gay", "biba"},
	}

	person["Sebastião"] = [2]string{"biba", "homo"}

	fmt.Println(person["Jorge"][1])
}

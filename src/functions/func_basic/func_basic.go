package main

import "fmt"

func sayHello(s string) (string, bool) {
	greetings := fmt.Sprintf("Hello, %s!", s)
	allowed_to_show := true

	return greetings, allowed_to_show
}

func main() {
	greetings, show := sayHello("Guilherme")

	if show {
		fmt.Println(greetings)
	} else {
		fmt.Println("Not allowed to display greetings")
	}
}

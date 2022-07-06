package main

import "fmt"

var myString string

func changeMyString(s *string, new_color string) {
	*s = new_color
}

func main() {
	myString = "Red"
	fmt.Println(myString)

	changeMyString(&myString, "Violet")
	fmt.Println(myString)
}

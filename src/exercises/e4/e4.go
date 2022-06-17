package main

import "fmt"

type str string

var gui str = "Guilherme"

func main() {
	name := fmt.Sprintf(`%v Antonio Massoqueto`, gui)
	fmt.Println(name)
}

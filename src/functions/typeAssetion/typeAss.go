package main

import "fmt"

type feline struct {
	specie      string
	is_domestic bool
	fountAt     string
}

type lion struct {
	feline
	weight uint16
}

type cat struct {
	feline
	name string
}

func (c cat) speak() {
	fmt.Println("Cat:", c.specie, "Found at:", c.fountAt)
}

func (l lion) speak() {
	fmt.Println("Lion:", l.specie, "Found at:", l.fountAt)
}

type animal interface {
	speak()
}

func assertType(t interface{}) string {
	switch t.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case cat:
		return "cat"
	case lion:
		return "lion"
	default:
		return "unknown"
	}
}

func main() {
	tom := cat{
		feline: feline{
			specie:      "Felis catus",
			fountAt:     "City",
			is_domestic: true,
		},
		name: "Tom",
	}

	mufasa := lion{
		feline: feline{
			specie:      "Panthera leo",
			fountAt:     "Savanna",
			is_domestic: false,
		},
		weight: 1_000,
	}

	mufasa.speak()
	tom.speak()
}

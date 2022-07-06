package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsPerson bool   `json:"is_person"`
}

func main() {
	persons := []person{}

	p1 := person{
		Name: "Jordi Polla",
		Age:  24,
	}

	p2 := person{
		Name: "Faustão",
		Age:  75,
	}

	persons = append(persons, p1, p2)

	fmt.Println(persons)

	newJason, error := json.MarshalIndent(persons, "", " ")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("%s", newJason)
}

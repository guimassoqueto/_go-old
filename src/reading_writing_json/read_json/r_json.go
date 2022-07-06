package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsPerson bool   `json:"is_person"`
}

func main() {
	/* Read Json */
	str_person := `
	[
		{"name":"Gaymen","age":100,"is_person":true},
		{"name":"NotGaymen","age":101,"is_person":false},
		{"name":"Guilherme","age":152,"is_person":true}
	]`
	person_struct := []person{}
	json.Unmarshal([]byte(str_person), &person_struct)

	for _, person := range person_struct {
		if person.IsPerson {
			fmt.Println(person)
		}
	}
}

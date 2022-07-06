package main

import (
	"errors"
)

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Cannot divide by zero.")
	}

	result = x / y
	return result, nil
}

func sum(a, b float32) (float32, error) {
	result := a + b
	return result, nil
}

func sub(a, b float32) (float32, error) {
	result := a - b
	return result, nil
}

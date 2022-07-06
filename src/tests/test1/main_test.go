package main

import (
	"testing"
)

/********** Inefficient ***************/
// func TestDivideSimple(t *testing.T) {
// 	_, err := divide(10.0, 10)

// 	if err != nil {
// 		t.Error("Error")
// 	}
// }

// func TestDivideByZero(t *testing.T) {
// 	_, err := divide(10.0, 0)

// 	if err == nil {
// 		t.Error("Should not throw error.")
// 	}
// }

// func TestDivideByNegative(t *testing.T) {
// 	_, err := divide(10.0, -10)

// 	if err != nil {
// 		t.Error("Error")
// 	}
// }

/************ Efficient ********************/

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-division", 10, 5, 2, false},
	{"divide-by-zero", 10, 0, 0, true},
	{"divide-by-negative", 10, -5, -2, false},
}

func TestDivide(t *testing.T) {
	for _, tt := range tests {
		divideResult, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Errorf("Expected error but not get one.")
			}
		} else {
			if err != nil {
				t.Errorf("Shoud not get an error but get one.")
			}
		}

		if tt.expected != divideResult {
			t.Errorf("Expected %v but got %v", tt.expected, divideResult)
		}
	}
}

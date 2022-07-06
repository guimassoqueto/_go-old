package main

import (
	"fmt"
	"math"
)

type circle struct {
	radious float64
}

type square struct {
	width  float64
	height float64
}

type form interface {
	area()
}

func (c circle) area() float64 {
	circle_area := math.Pi * math.Pow(c.radious, 2)
	return circle_area
}

func (s square) area() float64 {
	square_area := s.height * s.width
	return square_area
}

func main() {
	c := circle{
		radious: 3.23,
	}

	s := square{
		width:  10,
		height: 20,
	}

	fmt.Println("Circle:", c.area())
	fmt.Println("Square:", s.area())
}

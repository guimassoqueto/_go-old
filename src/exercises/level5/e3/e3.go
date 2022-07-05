package main

import "fmt"

type vehicle struct {
	door  uint8
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	f250 := truck{
		vehicle: vehicle{
			door:  2,
			color: "green",
		},
		fourWheel: false,
	}

	corolla := sedan{
		vehicle: vehicle{
			door:  4,
			color: "purple",
		},
		luxury: false,
	}

	fmt.Println(f250.color, f250.door, f250.fourWheel)
	fmt.Println(corolla.color, corolla.door, corolla.luxury)
}

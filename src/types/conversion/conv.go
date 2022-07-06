package main

import "fmt"

const f float32 = 0.255
const i uint8 = 255

func main() {
	num := float32(i)

	fmt.Printf("%T", num)
}

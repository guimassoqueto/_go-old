package main

import "fmt"

func incrementor(num int) func() int {
	return func() int {
		num++
		return num
	}
}

func main() {
	a := incrementor(100)
	b := incrementor(0)

	a()
	a()
	a()
	a()

	final_a := a()
	fmt.Println(final_a)

	for i := 0; i <= 10; i++ {
		b()
	}
	final_b := b()
	fmt.Println(final_b)

}

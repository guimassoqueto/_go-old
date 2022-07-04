package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		for j := 1; j <= 5; j++ {
			if j%2 == 0 {
				continue
			}
			fmt.Printf("%d.%d\n", i, j)
		}
	}
}

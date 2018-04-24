package main

import (
	"fmt"
)

// find the product of the three integers for which a + b + c = 1000 and a^2 + b^2 = c^2

func findPyth() int {
	for a := 1; a < 332; a++ {
		for b := a; b < 499; b++ {
			c := 1000 - a - b
			if (a*a + b*b) == (c * c) {
				product := a * b * c
				return product
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(findPyth())
}

package main

// Create an algorithm that will calculate the prime factors of a number.

import (
	"fmt"
	"math"
)

func findPrimeFactors(n int) {
	fmt.Println(n)
	// Remove even values
	for n%2 == 0 {
		fmt.Println(2)
		n = n / 2
	}

	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		for n%i == 0 {
			fmt.Println(i)
			n = n / i
		}
	}
	fmt.Println(n)
}

func main() {
	findPrimeFactors(28)
	findPrimeFactors(17)
	findPrimeFactors(126)
	findPrimeFactors(513)
}

package main

// Create an algorithm that will calculate the prime factors of a number.

import (
	"fmt"
	"math"
)

func findPrimeFactors(n int) {
	fmt.Println(n)

	// Account for zero
	if n == 0 {
		return
	}
	// Account for negative integers
	if n < 0 {
		n = int(math.Abs(float64(n)))
	}
	// Setup an output array
	var arr []int
	// Remove even values
	for n%2 == 0 {
		arr = append(arr, 2)
		n = n / 2
	}

	// cycle through possible odd factors
	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		for n%i == 0 {
			arr = append(arr, i)
			n = n / i
		}
	}
	if n > 2 {
		arr = append(arr, n)
	}
	fmt.Println(arr)
}

func main() {
	findPrimeFactors(28)
	findPrimeFactors(17)
	findPrimeFactors(126)
	findPrimeFactors(513)
	findPrimeFactors(5000)
	findPrimeFactors(-12)
	findPrimeFactors(0)
}

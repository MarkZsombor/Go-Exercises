package main

// Determing if a number is prime based on number of prime factors

import (
	"fmt"
	"math"
)

func findPrimeFactors(n int) (l int) {

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
	l = len(arr)
	return l
}

func isPrime(x int) (y bool) {
	n := findPrimeFactors(x)
	return n == 1
}

func main() {
	fmt.Println("Is 17 prime?")
	fmt.Println(isPrime(17))
	fmt.Println("Is 18 prime?")
	fmt.Println(isPrime(18))

}

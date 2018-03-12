package main

// Determing if a number is prime based on number of prime factors

import (
	"flag"
	"fmt"
	"math"
)

var args int

func init() {
	flag.IntVar(&args, "num", 0, "an int ")
}

func findPrimeFactors(n int) (a []int) {

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
	return arr
}

func isPrime(x int) (y bool) {
	n := findPrimeFactors(x)
	return len(n) == 1
}

func displayPrimeFactors(x int) {
	n := findPrimeFactors(x)
	if isPrime(x) || len(n) == 0 {
		fmt.Printf("%v is a Prime Number\n", x)
	} else {
		fmt.Printf("The prime factors of %v are:", x)
		fmt.Printf("%v\n", n)
	}

}

func main() {
	flag.Parse()
	if args != 0 {
		displayPrimeFactors(args)
	}

	displayPrimeFactors(1)
	displayPrimeFactors(67)
	displayPrimeFactors(6759)
	displayPrimeFactors(0)
	displayPrimeFactors(195128)
	displayPrimeFactors(9511)
	displayPrimeFactors(951)
	displayPrimeFactors(-64)
}

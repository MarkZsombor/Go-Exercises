package main

// Create an algorithm that will calculate the prime factors of a number.

import (
	"flag"
	"fmt"
	"math"
)

var args int

func init() {
	flag.IntVar(&args, "num", 0, "an int ")
}

func findPrimeFactors(n int) (arr []int) {

	// Account for zero
	if n == 0 {
		return
	}
	// Account for negative integers
	if n < 0 {
		n = int(math.Abs(float64(n)))
	}
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

func main() {
	fmt.Println(27, findPrimeFactors(27))
	fmt.Println(28, findPrimeFactors(28))
	fmt.Println(17, findPrimeFactors(17))
	fmt.Println(127, findPrimeFactors(127))
	fmt.Println(126, findPrimeFactors(126))
	fmt.Println(513, findPrimeFactors(513))
	fmt.Println(-66, findPrimeFactors(-66))
	fmt.Println(0, findPrimeFactors(0))
	fmt.Println(1, findPrimeFactors(1))

	flag.Parse()
	fmt.Println("From CLI: ", args, findPrimeFactors(args))

}

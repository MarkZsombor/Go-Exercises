package main

import (
	"fmt"
)

//what is the smallest positive number that is evenly divisible by all numbers 1-20

// Values given, when multiplied out, is too big by a factor of 3 with j*j!=i, too small by almost factor of 10 without
// Probably the wrong approach in general
func uniqueFactors(N int) (factors []int) {
	for i := N; i > 0; i-- {
		hasFactors := false
		for j := 2; j < i; j++ {
			if i%j == 0 { //&& j*j != i {
				hasFactors = true
			}
		}
		if !hasFactors {
			factors = append(factors, i)
		}
	}
	return
}

func multiplyArray(arr []int) (product int) {
	product = 1
	for i := 0; i < len(arr); i++ {
		product *= arr[i]
	}
	return
}

// This works, but WOW its brute forcing the issue and takes many seconds to complete
// Look into finding a better algorithm
func findDivisibleNumber(N int) int {
	testValue := N + 1
	for {
		isDivisible := true
		for i := 1; i <= N; i++ {
			if testValue%i != 0 {
				isDivisible = false
			}
		}
		if isDivisible {
			return testValue
		}
		testValue++
	}
}

func main() {
	// fmt.Println(findDivisibleNumber(20)) //232792560
	uf := uniqueFactors(10)
	fmt.Print(uf)
	fmt.Println(multiplyArray(uf))
}

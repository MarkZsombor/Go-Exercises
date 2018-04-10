package main

import (
	"fmt"
)

// Find the difference between the sum of the squares of first hundred and the square of the sum of said numbers

func sumOfSquares(num int) (SoSq int) {
	for i := 1; i <= num; i++ {
		SoSq += i * i
	}
	return
}

func squareOfSums(num int) (SqoS int) {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	SqoS = sum * sum
	return
}

func difference(SqoS, SoSq int) (diff int) {
	diff = SqoS - SoSq
	return
}

func main() {
	fmt.Println(sumOfSquares(10))
	fmt.Println(squareOfSums(10))
	fmt.Println(difference(squareOfSums(10), sumOfSquares(10)))
	fmt.Println(difference(squareOfSums(100), sumOfSquares(100)))

}

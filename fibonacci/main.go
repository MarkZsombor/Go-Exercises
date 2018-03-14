package main

import (
	"fmt"
)

// Stores all values in an array which is used to calc next value
func fibonacciArray(n int) (fibSeq []int) {
	fibSeq = append(fibSeq, 0)
	fibSeq = append(fibSeq, 1)
	for i := 1; i < n; i++ {
		fibSeq = append(fibSeq, (fibSeq[i] + fibSeq[i-1]))
	}
	return
}

// Stores only variables needed for next calculation
func fibonacci(n int) (sum int) {
	a := 0
	b := 1
	for i := 1; i < n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return
}

func main() {

	fmt.Println(fibonacci(4))
	fmt.Println(fibonacciArray(4))

	fmt.Println(fibonacci(10))
	fmt.Println(fibonacciArray(10))

	fmt.Println(fibonacci(20))
	fmt.Println(fibonacciArray(20))

}

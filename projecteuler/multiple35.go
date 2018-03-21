package main

import (
	"fmt"
)

// find the sum of all multiples of 3 and 5 less than 1000

func sumMultiples(N int) (sum int) {
	for i := 1; i < N; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return
}

func main() {
	fmt.Println(sumMultiples(10))
	fmt.Println(sumMultiples(100))
	fmt.Println(sumMultiples(1000))

}

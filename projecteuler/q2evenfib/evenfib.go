package main

import (
	"fmt"
)

// Using a fibonacci sequence starting with 1,2 sum the even value terms that are less than 4M
func fibonacci(n int) int {
	if n <= 2 {
		return n
	} else {
		return (fibonacci(n-1) + fibonacci(n-2))
	}
}

func fibSum(N int) (sum int) {
	var i int
	for fibonacci(i) < N {
		if fibonacci(i)%2 == 0 {
			sum += fibonacci(i)
		}
		i++
	}
	return
}

func main() {
	fmt.Println(fibSum(4000000))
}

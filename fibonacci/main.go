package main

import (
	"fmt"
)

func fibonacci(n int) {
	var fibSeq []int
	fibSeq = append(fibSeq, 0)
	fibSeq = append(fibSeq, 1)
	for i := 1; i < n; i++ {
		fibSeq = append(fibSeq, (fibSeq[i] + fibSeq[i-1]))
	}
	fmt.Println(fibSeq)
}

func main() {
	fibonacci(4)
	fibonacci(10)
	fibonacci(20)
}

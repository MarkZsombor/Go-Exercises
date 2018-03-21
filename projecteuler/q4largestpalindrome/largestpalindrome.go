package main

import (
	"fmt"
)

// Find the largest palindrome number that is the product of two 3-didgit numbers
func reverseNumber(n int) (rev int) {
	for {
		decimalValue := n % 10
		rev = rev*10 + decimalValue
		n /= 10

		if n == 0 {
			break
		}
	}
	return
}

func isPalindrome(n int) bool {
	return n == reverseNumber(n)
}

func findLargestPalindrome() (large int) {
	large = 0
	for i := 900; i < 1000; i++ {
		for j := 900; j < 1000; j++ {
			product := i * j
			if isPalindrome(product) && product > large {
				large = product
			}
		}
	}
	return
}

func main() {
	fmt.Println(46589, reverseNumber(46589), isPalindrome(46589))
	fmt.Println(45654, reverseNumber(45654), isPalindrome(45654))
	fmt.Println(findLargestPalindrome())
}

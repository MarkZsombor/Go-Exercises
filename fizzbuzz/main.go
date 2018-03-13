package main

import (
	"fmt"
)

/* The classic, cycle through the numbers from 1-100, if a number is divisible by 3 print Fizz, if its divisible by 5 print Buzz, divisible by 15 print FizzBuzz, else print the number */
func fizzBuzz() {
	var output string
	for i := 1; i <= 100; i++ {
		output = ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if output == "" {
			fmt.Println(i)
		} else {
			fmt.Println(output)
		}
	}
}

func main() {
	fizzBuzz()
}

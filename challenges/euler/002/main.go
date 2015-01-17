package main

import (
	"fmt"
	"projects/internal/euler"
)

// Even Fibonacci numbers
// https://projecteuler.net/problem=2
func main() {
	fmt.Println("The solution is:", solve())
}

func solve() int {
	var sum int

	for _, val := range euler.GenFibUnder(4000000) {
		if val%2 == 0 {
			sum += val
		}
	}

	return sum
}

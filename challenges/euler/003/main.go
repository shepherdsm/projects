package main

import "fmt"

// Largest Prime Factor
// https://projecteuler.net/problem=3
func main() {
	fmt.Println("The solution to challenge 3:", solve())
}

func solve() int {
	start := 600851475143
	div := 3

	for div*div <= start {
		if start%div == 0 {
			start /= div
		} else {
			div += 2
		}
	}

	return start
}

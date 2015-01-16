package main

import "fmt"

// Multiples of 3 and 5.
// https://projecteuler.net/problem=1
func main() {
	fmt.Println("The solution to challenge 1 is:", solve())
}

func solve() int {
	var sum int

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

package main

import (
	"fmt"
	"math"
)

// Prime Palindrome
// https://www.codeeval.com/open_challenges/3/
func main() {
	for k := 999; k >= 100; k-- {
		if isPalindrome(k) && isPrime(k) {
			fmt.Println(k)
			break
		}
	}
}

// Assumes 3 digit numbers.
func isPalindrome(n int) bool {
	return n%10 == n/100
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	upper := int(math.Ceil(math.Sqrt(float64(n))))
	div := 3

	for div <= upper {
		if n%div == 0 {
			return false
		}

		div += 2
	}

	return true
}

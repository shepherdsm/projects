package main

import (
	"fmt"
	"math"
)

// Sum Of Primes
// https://www.codeeval.com/open_challenges/4/
func main() {
	sum := 2
	num := 3
	count := 1

	for count < 1000 {
		if isPrime(num) {
			sum += num
			count++
		}

		num += 2
	}

	fmt.Println(sum)
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

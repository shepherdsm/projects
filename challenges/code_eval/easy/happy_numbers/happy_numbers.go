package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Happy Numbers
// https://www.codeeval.com/open_challenges/39
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(happy(num))
	}
}

func happy(n int) int {
	seen := make(map[int]bool)

	for {
		if _, ok := seen[n]; !ok {
			seen[n] = true
			n = digitSquareSum(n)

			if n == 1 {
				return 1
			}
		} else {
			return 0
		}
	}
}

func digitSquareSum(n int) int {
	sum := 0

	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}

	return sum
}

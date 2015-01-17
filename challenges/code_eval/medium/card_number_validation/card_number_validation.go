package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Card Number Validation
// https://www.codeeval.com/open_challenges/172/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if validateCC(scanner.Text()) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}

func validateCC(s string) bool {
	double := false
	var sum int

	for k := len(s) - 1; k >= 0; k -= 1 {
		if char := rune(s[k]); unicode.IsNumber(char) {
			num := int(char - '0')

			if double {
				num *= 2

				if num >= 10 {
					num -= 9
				}
			}

			double = !double
			sum += num
		}
	}

	return sum%10 == 0
}

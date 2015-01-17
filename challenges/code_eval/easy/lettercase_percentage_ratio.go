package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Lettercase Percentage Ratio
// https://www.codeeval.com/open_challenges/147
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		upper, lower := ratio(scanner.Text())
		fmt.Printf("lowercase: %2.2f uppercase: %2.2f\n", lower, upper)
	}
}

func ratio(s string) (float64, float64) {
	var (
		upper float64
		lower float64
	)

	for _, c := range s {
		if unicode.IsUpper(c) {
			upper++
		} else {
			lower++
		}
	}

	total := upper + lower

	return 100 * upper / total, 100 * lower / total
}

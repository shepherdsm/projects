package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Word To Digit
// https://www.codeeval.com/open_challenges/104/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	WORDMAP := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ";")

		for _, val := range items {
			fmt.Print(WORDMAP[val])
		}

		fmt.Println()
	}
}

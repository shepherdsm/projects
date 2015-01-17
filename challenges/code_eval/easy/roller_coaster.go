package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// Roller Coaster
// https://www.codeeval.com/open_challenges/156/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		roller(scanner.Text())
	}
}

func roller(s string) {
	up := true

	for _, ch := range s {
		if unicode.IsLetter(ch) {
			if up {
				fmt.Print(string(unicode.ToUpper(ch)))
			} else {
				fmt.Print(string(unicode.ToLower(ch)))
			}

			up = !up
		} else {
			fmt.Print(string(ch))
		}
	}

	fmt.Println()
}

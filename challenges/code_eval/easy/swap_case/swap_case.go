package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, ch := range scanner.Text() {
			if unicode.IsUpper(ch) {
				fmt.Print(string(unicode.ToLower(ch)))
			} else {
				fmt.Print(string(unicode.ToUpper(ch)))
			}
		}

		fmt.Println()
	}
}

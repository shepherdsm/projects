package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// First Non-Repeated Character
// https://www.codeeval.com/open_challenges/12/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(firstNonRepChar(scanner.Text()))
	}
}

// Only works for ASCII characters.
func firstNonRepChar(s string) string {
	for _, ch := range s {
		sep := string(ch)

		if strings.Count(s, sep) == 1 {
			return sep
		}
	}

	return ""
}

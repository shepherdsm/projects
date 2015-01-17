package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Longest Word
// https://www.codeeval.com/open_challenges/111/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(longestWord(strings.Split(scanner.Text(), " ")))
	}
}

func longestWord(s []string) string {
	longest := s[0]

	for _, str := range s[1:] {
		if len(str) > len(longest) {
			longest = str
		}
	}

	return longest
}

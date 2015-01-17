package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Remove Characters
// https://www.codeeval.com/open_challenges/13/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ", ")

		mapper := make(map[rune]bool)
		for _, val := range items[1] {
			mapper[val] = true
		}

		drop := func(r rune) rune {
			if _, ok := mapper[r]; ok {
				return -1
			} else {
				return r
			}
		}

		fmt.Println(strings.Map(drop, items[0]))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Without Repetitions
// https://www.codeeval.com/open_challenges/173/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var prev rune
		for _, val := range scanner.Text() {
			if val == prev {
				continue
			}

			fmt.Printf("%c", val)
			prev = val
		}
		fmt.Println()
	}
}

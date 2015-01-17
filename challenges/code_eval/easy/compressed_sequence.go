package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Compressed Sequence
// https://www.codeeval.com/open_challenges/128/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		first := line[0]
		count := 1

		for _, val := range line[1:] {
			if first == val {
				count++
			} else {
				fmt.Printf("%d %s ", count, first)
				first = val
				count = 1
			}
		}

		fmt.Printf("%d %s\n", count, first)
	}
}

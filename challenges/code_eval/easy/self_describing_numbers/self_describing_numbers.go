package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Self Describing Numbers
// https://www.codeeval.com/open_challenges/40/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(isSelfDescribing(scanner.Text()))
	}
}

func isSelfDescribing(s string) int {
	for ind, val := range s {
		// Converts the val rune to its integer representation.
		num := int(val - '0')

		if strings.Count(s, strconv.Itoa(ind)) != num {
			return 0
		}
	}

	return 1
}

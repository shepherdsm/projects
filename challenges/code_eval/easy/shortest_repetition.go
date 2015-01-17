package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(shortestRep(scanner.Text()))
	}
}

func shortestRep(s string) int {
	strLen := len(s)

	for k := 0; k < strLen; k++ {
		sub := s[:k]
		count := strings.Count(s, s[:k])
		subLen := len(sub)

		if subLen*count == strLen {
			return subLen
		}
	}

	return strLen
}

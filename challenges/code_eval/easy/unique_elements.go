package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Unique Elements
// https://www.codeeval.com/open_challenges/29
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")
		seen := make(map[string]bool)
		var results []string

		for _, n := range nums {
			if _, ok := seen[n]; !ok {
				results = append(results, n)
				seen[n] = true
			}
		}

		fmt.Println(strings.Join(results, ","))
	}
}

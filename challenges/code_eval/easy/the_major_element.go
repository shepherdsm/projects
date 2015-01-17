package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The Major Element
// https://www.codeeval.com/open_challenges/132/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(majorElement(strings.Split(scanner.Text(), ",")))
	}
}

func majorElement(items []string) string {
	major := len(items) / 2
	nums := make(map[string]int)

	for _, val := range items {
		nums[val]++
	}

	for key, val := range nums {
		if val > major {
			return key
		}
	}

	return "None"
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

// Beautiful Strings
// https://www.codeeval.com/open_challenges/83/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(beautiful(scanner.Text()))
	}
}

func beautiful(s string) int {
	var (
		sum  int
		nums []int
	)
	seen := make(map[rune]bool)
	s = strings.ToLower(s)

	for _, c := range s {
		if _, ok := seen[c]; !ok {
			seen[c] = true

			if unicode.IsLetter(c) {
				nums = append(nums, strings.Count(s, string(c)))
			}
		}
	}

	count := 26
	sort.Ints(nums)
	for k := len(nums) - 1; k >= 0; k-- {
		sum += nums[k] * count
		count--
	}

	return sum
}

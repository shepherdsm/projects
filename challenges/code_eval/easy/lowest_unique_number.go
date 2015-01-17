package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Lowest Unique Number
// https://www.codeeval.com/open_challenges/103/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var nums []int

		for _, val := range strings.Split(scanner.Text(), " ") {
			tmp, _ := strconv.Atoi(val)
			nums = append(nums, tmp)
		}

		fmt.Println(lowestUnique(nums))
	}
}

func lowestUnique(nums []int) int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	cur := 0
	var count int
	for _, val := range sorted {
		if cur != val {
			if count == 1 {
				break
			} else if count > 1 {
				count = 0
			}
		}

		cur = val
		count++
	}

	if count > 1 {
		cur = 0
	}

	if cur != 0 {
		for ind, val := range nums {
			if cur == val {
				return ind + 1
			}
		}
	}

	return 0
}

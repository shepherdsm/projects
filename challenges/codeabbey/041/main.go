package main

import (
	"bufio"
	"fmt"
	"os"
)

// Median of Three
// http://www.codeabbey.com/index/task_view/median-of-three
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		a int
		b int
		c int
	)
	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &a, &b, &c)

		fmt.Printf("%d ", max(min(a, b), min(max(a, b), c)))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

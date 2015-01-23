package main

import (
	"bufio"
	"fmt"
	"os"
)

// Minimum of Three
// http://www.codeabbey.com/index/task_view/min-of-three
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		a int
		b int
		c int
	)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d %d", &a, &b, &c)
		fmt.Printf("%d ", min(a, min(b, c)))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

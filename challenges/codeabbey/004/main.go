package main

import (
	"bufio"
	"fmt"
	"os"
)

// Minimum of Two
// http://www.codeabbey.com/index/task_view/min-of-two
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		a int
		b int
	)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)

		if a < b {
			fmt.Printf("%d ", a)
		} else {
			fmt.Printf("%d ", b)
		}
	}
}

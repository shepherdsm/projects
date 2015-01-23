package main

import (
	"bufio"
	"fmt"
	"os"
)

// Sums in Loop
// http://www.codeabbey.com/index/task_view/sums-in-loop
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		a int
		b int
	)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
		fmt.Printf("%d ", a+b)
	}
}

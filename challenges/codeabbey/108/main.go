package main

import (
	"bufio"
	"fmt"
	"os"
)

// Star Medals
// http://www.codeabbey.com/index/task_view/star-medals
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		n int
		t int
	)

	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &n, &t)
		fmt.Printf("%d ", (n * (t - 1)))
	}
}

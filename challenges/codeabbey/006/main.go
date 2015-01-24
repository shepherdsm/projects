package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Rounding
// http://www.codeabbey.com/index/task_view/rounding
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var (
		top    float64
		bottom float64
	)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%f %f", &top, &bottom)
		fmt.Printf("%d ", round(top/bottom))
	}
}

func round(num float64) int {
	return int(math.Floor(num + 0.5))
}

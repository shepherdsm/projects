package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Calculate Distance
// https://www.codeeval.com/open_challenges/99/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x1, y1, x2, y2 int
		fmt.Sscanf(scanner.Text(), "(%d, %d) (%d, %d)", &x1, &y1, &x2, &y2)
		fmt.Println(distance(x1, y1, x2, y2))
	}
}

// Problem guarantees pairs produce integer values
func distance(x1, y1, x2, y2 int) int {
	sum1, sum2 := float64(x2-x1), float64(y2-y1)
	return int(math.Sqrt((sum1*sum1 + sum2*sum2)))
}

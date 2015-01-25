package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Dice Rolling
// http://www.codeabbey.com/index/task_view/dice-rolling
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var num float64

	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &num)

		fmt.Printf("%d ", int(math.Floor(num*6.0))+1)
	}
}

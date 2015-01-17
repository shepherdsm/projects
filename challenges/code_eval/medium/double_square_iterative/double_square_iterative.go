package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Double Square
// https://www.codeeval.com/open_challenges/33/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(numWays(float64(num)))
	}
}

func numWays(n float64) int {
	count := 0
	upper := math.Floor(math.Sqrt(n))

	for i := 0.0; i <= upper; i++ {
		x := math.Pow(i, 2)
		y := n - x

		if y < x {
			break
		}

		if val := math.Sqrt(y); val == math.Floor(val) {
			count++
		}
	}

	return count
}

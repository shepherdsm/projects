package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

// Double Square
// https://www.codeeval.com/open_challenges/33/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	cases, _ := strconv.Atoi(scanner.Text())

	results := make([]int, cases)
	wg.Add(cases)

	numWays := func(n float64, index int) {
		defer wg.Done()

		upper := math.Floor(math.Sqrt(n))

		for i := 0.0; i <= upper; i++ {
			x := math.Pow(i, 2)
			y := n - x

			if y < x {
				break
			}

			if val := math.Sqrt(y); val == math.Floor(val) {
				results[index] += 1
			}
		}
	}

	count := 0

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		go numWays(float64(num), count)
		count++
	}

	wg.Wait()

	for _, val := range results {
		fmt.Println(val)
	}
}

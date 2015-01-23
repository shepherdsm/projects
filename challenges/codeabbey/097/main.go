package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Girls and Pigs
// http://www.codeabbey.com/index/task_view/girls-and-pigs
func main() {
	var breasts, legs int

	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d", &legs, &breasts)

		fmt.Print(getSolutions(breasts, legs), " ")
	}
}

func getSolutions(breasts, legs int) int {
	numPigs := func(pigBreasts int) float64 {
		return float64(breasts-legs) / float64(pigBreasts-4)
	}

	var solutions int
	pigBreasts := 6

	for {
		pigs := numPigs(pigBreasts)
		if pigs < 1 {
			break
		} else if math.Floor(pigs) != pigs {
			pigBreasts += 2
			continue
		}

		if girls := (breasts - pigBreasts*int(pigs)) / 2; girls > 0 {
			solutions++
		}

		pigBreasts += 2
	}

	return solutions
}

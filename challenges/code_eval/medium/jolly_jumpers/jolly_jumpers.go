package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Jolly Jumpers
// https://www.codeeval.com/open_challenges/43/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		nums := make([]int, len(items)-1)
		compare := make(map[int]bool)

		for ind, val := range items[1:] {
			tmp, _ := strconv.Atoi(val)
			nums[ind] = tmp

			if ind > 0 {
				compare[ind] = false
			}
		}

		jolly := true
		prev := nums[0]
		for _, val := range nums[1:] {
			dif := val - prev
			prev = val

			if dif < 0 {
				dif *= -1
			}

			if _, ok := compare[dif]; ok {
				compare[dif] = true
			} else {
				jolly = false
				break
			}
		}

		for _, val := range compare {
			jolly = jolly && val
		}

		if jolly {
			fmt.Println("Jolly")
		} else {
			fmt.Println("Not jolly")
		}
	}
}

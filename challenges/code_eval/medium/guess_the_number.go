package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Guess The Number
// https://www.codeeval.com/open_challenges/170/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		ub, _ := strconv.Atoi(items[0])
		items = items[1:]

		arr := make([]int, ub)
		for k := 0; k < ub; k++ {
			arr[k] = k
		}

		fmt.Println(binSearch(0, ub, arr, items))
	}
}

func binSearch(lower, upper int, arr []int, dir []string) int {
	tmp := upper - lower
	if tmp%2 != 0 {
		tmp++
	}
	mid := lower + (tmp / 2)
	d := dir[0]

	if d == "Yay!" {
		return arr[mid]
	} else {
		if d == "Lower" {
			return binSearch(lower, mid-1, arr, dir[1:])
		} else {
			return binSearch(mid+1, upper, arr, dir[1:])
		}
	}
}

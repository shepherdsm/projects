package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Multiples Of A Number
// https://www.codeeval.com/open_challenges/18/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(nums[0])
		n, _ := strconv.Atoi(nums[1])

		fmt.Println(multiple(x, n))
	}
}

func multiple(bound, start int) int {
	res := start

	for res < bound {
		res += start
	}

	return res
}

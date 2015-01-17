package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Split The Number
// https://www.codeeval.com/open_challenges/131/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		fmt.Println(processLine(items[0], items[1]))
	}
}

func processLine(num string, pattern string) int {
	ind := strings.Index(pattern, "+")
	op := func(x, y int) int { return x + y }

	if ind == -1 {
		ind = strings.Index(pattern, "-")
		op = func(x, y int) int { return x - y }
	}

	left, _ := strconv.Atoi(num[:ind])
	right, _ := strconv.Atoi(num[ind:])

	return op(left, right)
}

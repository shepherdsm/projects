package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Fibonacci Series
// https://www.codeeval.com/open_challenges/22/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(fibonacci(num))
	}
}

func fibonacci(n int) int {
	prev, cur := 0, 1

	for i := 0; i < n; i++ {
		cur, prev = cur+prev, cur
	}

	return prev
}

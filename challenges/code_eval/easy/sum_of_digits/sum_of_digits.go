package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Sum of Digits
// https://www.codeeval.com/open_challenges/21/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(digitSum(num))
	}
}

func digitSum(n int) int {
	var sum int

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

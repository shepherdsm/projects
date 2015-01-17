package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Number Of Ones
// https://www.codeeval.com/open_challenges/16/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(countOnes(num))
	}
}

func countOnes(n int) int {
	var count int

	for n > 0 {
		count += n & 1
		n >>= 1
	}

	return count
}

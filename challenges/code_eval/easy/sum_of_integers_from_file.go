package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Sum of Integers From File
// https://www.codeeval.com/open_challenges/24/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		sum += num
	}

	fmt.Println(sum)
}

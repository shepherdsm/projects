package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Hex To Decimal
// https://www.codeeval.com/open_challenges/67/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.ParseInt(scanner.Text(), 16, 0)
		fmt.Println(num)
	}
}

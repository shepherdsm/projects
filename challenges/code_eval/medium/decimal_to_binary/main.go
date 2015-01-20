package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Decimal To Binary
// https://www.codeeval.com/open_challenges/27/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(strconv.FormatInt(int64(num), 2))
	}
}

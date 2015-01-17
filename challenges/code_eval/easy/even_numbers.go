package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Even Numbers
// https://www.codeeval.com/open_challenges/100/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())

		if num%2 == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}

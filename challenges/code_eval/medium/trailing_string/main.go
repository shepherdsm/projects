package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Trailing String
// https://www.codeeval.com/open_challenges/32/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ",")

		if strings.HasSuffix(lines[0], lines[1]) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}

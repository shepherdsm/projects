package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Stack Implementation
// https://www.codeeval.com/open_challenges/9/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")

		for i := len(items) - 1; i >= 0; i -= 2 {
			fmt.Printf("%s ", items[i])
		}

		fmt.Println()
	}
}

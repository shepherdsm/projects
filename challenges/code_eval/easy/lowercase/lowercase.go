package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Lowercase
// https://www.codeeval.com/open_challenges/20/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(strings.ToLower(scanner.Text()))
	}
}

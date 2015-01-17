package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Capitalize Words
// https://www.codeeval.com/open_challenges/93/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(strings.Title(scanner.Text()))
	}
}

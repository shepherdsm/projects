package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Rightmost Char
// https://www.codeeval.com/open_challenges/31
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ",")
		fmt.Println(strings.LastIndex(items[0], items[1]))
	}
}

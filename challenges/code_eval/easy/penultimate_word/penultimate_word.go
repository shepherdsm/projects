package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Penultimate Word
// https://www.codeeval.com/open_challenges/92
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		fmt.Println(words[len(words)-2])
	}
}

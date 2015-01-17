package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Reverse Words
// https://www.codeeval.com/open_challenges/8/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reverse(scanner.Text())
	}
}

func reverse(s string) {
	words := strings.Split(s, " ")

	for k := len(words) - 1; k >= 0; k-- {
		fmt.Print(words[k])

		if k != 0 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
}

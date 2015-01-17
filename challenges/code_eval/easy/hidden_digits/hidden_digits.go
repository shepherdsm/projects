package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Hidden Digits
// https://www.codeeval.com/open_challenges/122/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if res := strings.Map(mapper, scanner.Text()); res != "" {
			fmt.Println(res)
		} else {
			fmt.Println("NONE")
		}
	}
}

func mapper(r rune) rune {
	switch {
	case r >= 'a' && r <= 'j':
		return r - 'a' + '0'
	case r >= '0' && r <= '9':
		return r
	}

	return -1
}

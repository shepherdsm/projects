package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Read More
// https://www.codeeval.com/open_challenges/167/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(prepLine(scanner.Text()))
	}
}

func prepLine(s string) string {
	if len(s) > 55 {
		s = s[:40]
		if ind := strings.LastIndex(s, " "); ind != -1 {
			s = s[:ind]
		}

		s += "... <Read More>"
	}

	return s
}

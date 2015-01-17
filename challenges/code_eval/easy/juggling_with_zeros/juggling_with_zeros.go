package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Juggling With Zeroes
// https://www.codeeval.com/open_challenges/149/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(parseLine(strings.Split(scanner.Text(), " ")))
	}
}

func parseLine(items []string) int {
	var num int

	for k := 0; k < len(items); k += 2 {
		num = parseValue(num, items[k], items[k+1])
	}

	return num
}

func parseValue(prev int, flag, amount string) int {
	if prev == 0 && flag == "0" {
		return 0
	}

	switch {
	case flag == "00":
		for k := 0; k < len(amount); k++ {
			prev <<= 1
			prev += 1
		}
	case flag == "0":
		prev <<= uint(len(amount))
	}

	return prev
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// N Mod M
// https://www.codeeval.com/open_challenges/62/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ",")
		n, _ := strconv.Atoi(items[0])
		m, _ := strconv.Atoi(items[1])

		fmt.Println(mod(n, m))
	}
}

func mod(n, m int) int {
	for n >= m {
		n -= m
	}

	return n
}

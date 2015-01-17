package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Mth to Last Element
// https://www.codeeval.com/open_challenges/10/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		count := len(items) - 1
		num, _ := strconv.Atoi(items[count])
		items = items[:count]

		if count >= num {
			fmt.Println(items[count-num])
		}
	}
}

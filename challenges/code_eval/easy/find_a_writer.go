package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Find A Writer
// https://www.codeeval.com/open_challenges/97/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(findWriter(scanner.Text()))
	}
}

func findWriter(code string) string {
	items := strings.Split(code, "| ")
	enc := items[0]
	items = strings.Split(items[1], " ")
	var res string

	for _, val := range items {
		num, _ := strconv.Atoi(val)
		res += string(enc[num-1])
	}

	return res
}

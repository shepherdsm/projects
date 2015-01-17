package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Bit Positions
// https://www.codeeval.com/open_challenges/19/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ",")
		fmt.Println(sameBits(items))
	}
}

func sameBits(items []string) string {
	num, _ := strconv.Atoi(items[0])
	bits := strconv.FormatInt(int64(num), 2)
	p1, _ := strconv.Atoi(items[1])
	p2, _ := strconv.Atoi(items[2])
	l := len(bits)

	if bits[l-p1] == bits[l-p2] {
		return "true"
	} else {
		return "false"
	}
}

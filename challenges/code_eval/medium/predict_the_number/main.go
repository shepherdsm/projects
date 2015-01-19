package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Predict The Number
// https://www.codeeval.com/open_challenges/125/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(getSeqN(num))
	}
}

func getOnes(n int) int {
	bin := strconv.FormatInt(int64(n), 2)
	return strings.Count(bin, "1")
}

func getSeqN(n int) int {
	return getOnes(n) % 3
}

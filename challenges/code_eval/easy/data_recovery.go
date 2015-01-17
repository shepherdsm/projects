package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Data Recovery
// https://www.codeeval.com/open_challenges/140/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ";")
		fmt.Println(rebuild(items[0], items[1]))
	}
}

func rebuild(s string, numString string) string {
	wordMap := make(map[int]string)
	words := strings.Split(s, " ")

	for ind, val := range strings.Split(numString, " ") {
		tmp, _ := strconv.Atoi(val)
		wordMap[tmp] = words[ind]
	}

	wordMap[-1] = words[len(words)-1]
	fixed := make([]string, len(words))

	for k := 1; k <= len(words); k++ {
		if val, ok := wordMap[k]; ok {
			fixed[k-1] = val
		} else {
			fixed[k-1] = wordMap[-1]
		}
	}

	return strings.Join(fixed, " ")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Simple Sorting
// https://www.codeeval.com/open_challenges/91/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%s\n", sortNums(strings.Split(scanner.Text(), " ")))
	}
}

func sortNums(items []string) string {
	var res []float64

	for _, val := range items {
		tmp, _ := strconv.ParseFloat(val, 64)
		res = append(res, tmp)
	}

	sort.Float64s(res)
	var tmp []string

	for _, val := range res {
		tmp = append(tmp, strconv.FormatFloat(val, 'f', 3, 64))
	}

	return strings.Join(tmp, " ")
}

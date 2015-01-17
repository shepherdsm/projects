package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Set Intersection
// https://www.codeeval.com/open_challenges/30/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(intersection(scanner.Text()))
	}
}

func intersection(s string) string {
	lists := strings.Split(s, ";")
	set1 := strings.Split(lists[0], ",")
	set2 := strings.Split(lists[1], ",")
	var results []string

	for _, v1 := range set1 {
		val1, _ := strconv.Atoi(v1)

		for _, v2 := range set2 {
			val2, _ := strconv.Atoi(v2)

			if val1 == val2 {
				results = append(results, v2)
			}
		}
	}

	return strings.Join(results, ",")
}

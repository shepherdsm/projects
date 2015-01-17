package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Multiply Lists
// https://www.codeeval.com/open_challenges/113/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lists := strings.Split(scanner.Text(), " | ")
		var results []string
		list1 := strings.Split(lists[0], " ")
		list2 := strings.Split(lists[1], " ")

		for ind, val := range list1 {
			num1, _ := strconv.Atoi(string(val))
			num2, _ := strconv.Atoi(string(list2[ind]))

			results = append(results, strconv.Itoa(num1*num2))
		}

		fmt.Println(strings.Join(results, " "))
	}
}

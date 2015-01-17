package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reverse Groups
// https://www.codeeval.com/open_challenges/71/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), ";")
		groups, _ := strconv.Atoi(items[1])
		nums := strings.Split(items[0], ",")

		fmt.Println(strings.Join(revGroups(nums, groups), ","))
	}
}

func revGroups(items []string, group int) []string {
	rev := make([]string, len(items), len(items))

	for k := 0; k < len(items); k += group {
		tmp := items[k:]

		if group <= len(tmp) {
			for i := range tmp[:group] {
				rev[k+i] = tmp[group-i-1]
			}
		} else {
			copy(rev[k:], tmp)
		}
	}

	return rev
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Mixed Content
// https://www.codeeval.com/open_challenges/115/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var str []string
		var nums []string

		for _, val := range strings.Split(scanner.Text(), ",") {
			if _, err := strconv.Atoi(val); err != nil {
				str = append(str, val)
			} else {
				nums = append(nums, val)
			}
		}

		var out string
		if len(str) > 0 {
			out += strings.Join(str, ",")
		}
		if len(nums) > 0 {
			if len(out) > 0 {
				out += "|"
			}
			out += strings.Join(nums, ",")
		}

		fmt.Println(out)
	}
}

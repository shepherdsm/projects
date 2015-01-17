package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fizz Buzz
// https://www.codeeval.com/open_challenges/1/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		fizz, _ := strconv.Atoi(items[0])
		buzz, _ := strconv.Atoi(items[1])
		count, _ := strconv.Atoi(items[2])

		for k := 1; k <= count; k++ {
			var out string

			if k%fizz == 0 {
				out += "F"
			}
			if k%buzz == 0 {
				out += "B"
			}
			if len(out) == 0 {
				out += strconv.Itoa(k)
			}
			if k != count {
				out += " "
			}

			fmt.Print(out)
		}

		fmt.Println()
	}
}

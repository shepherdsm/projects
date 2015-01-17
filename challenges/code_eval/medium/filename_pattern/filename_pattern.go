package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Filename Pattern
// https://www.codeeval.com/open_challenges/169/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")

		re := filePatToRegex(items[0])
		var res string

		for _, val := range items[1:] {
			if match := re.FindString(val); match != "" {
				res += fmt.Sprintf(" %s", match)
			}
		}

		if len(res) > 0 {
			fmt.Println(res[1:])
		} else {
			fmt.Println("-")
		}
	}
}

func filePatToRegex(fp string) *regexp.Regexp {
	fp = strings.Replace(fp, ".", `\.`, -1)
	fp = strings.Replace(fp, "?", ".", -1)
	fp = strings.Replace(fp, "*", ".*", -1)
	fp = "^" + fp + "$"

	regex, _ := regexp.Compile(fp)

	return regex
}

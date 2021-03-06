package main

import (
	"bufio"
	"fmt"
	"os"
)

// Big Digits
// https://www.codeeval.com/open-challenges/163/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	numMap := map[rune][]string{
		'0': {"-**--", "*--*-", "*--*-", "*--*-", "-**--", "-----"},
		'1': {"--*--", "-**--", "--*--", "--*--", "-***-", "-----"},
		'2': {"***--", "---*-", "-**--", "*----", "****-", "-----"},
		'3': {"***--", "---*-", "-**--", "---*-", "***--", "-----"},
		'4': {"-*---", "*--*-", "****-", "---*-", "---*-", "-----"},
		'5': {"****-", "*----", "***--", "---*-", "***--", "-----"},
		'6': {"-**--", "*----", "***--", "*--*-", "-**--", "-----"},
		'7': {"****-", "---*-", "--*--", "-*---", "-*---", "-----"},
		'8': {"-**--", "*--*-", "-**--", "*--*-", "-**--", "-----"},
		'9': {"-**--", "*--*-", "-***-", "---*-", "-**--", "-----"},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for k := 0; k < 6; k++ {
			for _, ch := range scanner.Text() {
				if val, ok := numMap[ch]; ok {
					fmt.Print(val[k])
				}
			}

			fmt.Println()
		}
	}
}

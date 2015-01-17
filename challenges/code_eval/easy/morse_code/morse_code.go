package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Morse Code
// https://www.codeeval.com/open_challenges/116/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(morsify(scanner.Text()))
	}
}

func morsify(s string) string {
	mapping := map[string]rune{
		".-":    'A',
		"-...":  'B',
		"-.-.":  'C',
		"-..":   'D',
		".":     'E',
		"..-.":  'F',
		"--.":   'G',
		"....":  'H',
		"..":    'I',
		".---":  'J',
		"-.-":   'K',
		".-..":  'L',
		"--":    'M',
		"-.":    'N',
		"---":   'O',
		".--.":  'P',
		"--.-":  'Q',
		".-.":   'R',
		"...":   'S',
		"-":     'T',
		"..-":   'U',
		"...-":  'V',
		".--":   'W',
		"-..-":  'X',
		"-.--":  'Y',
		"--..":  'Z',
		".----": '1',
		"..---": '2',
		"...--": '3',
		"....-": '4',
		".....": '5',
		"-....": '6',
		"--...": '7',
		"---..": '8',
		"----.": '9',
		"-----": '0',
	}

	var out []string

	for _, word := range strings.Split(s, "  ") {
		var w string
		for _, mc := range strings.Split(word, " ") {
			w += string(mapping[mc])
		}

		out = append(out, w)
	}

	return strings.Join(out, " ")
}

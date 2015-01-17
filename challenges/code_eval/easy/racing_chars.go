package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Racing Chars
// https://www.codeeval.com/open_challenges/136/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var out string
	prev := -1
	for scanner.Scan() {
		out, prev = processLine(scanner.Text(), prev)
		fmt.Println(out)
	}
}

func processLine(s string, prev int) (string, int) {
	track := strings.Index(s, "C")

	if track == -1 {
		track = strings.Index(s, "_")
	}

	var dir string
	if track == prev || prev == -1 {
		dir = "|"
	} else if track > prev {
		dir = "\\"
	} else {
		dir = "/"
	}

	var out string
	if track == 0 {
		out = dir + s[track+1:]
	} else if track == len(s)-1 {
		out = s[:track-1] + dir
	} else {
		out = s[:track] + dir + s[track+1:]
	}

	return out, track
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Slang Flavor
// https://www.codeeval.com/open_challenges/174/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	replace := false
	slangIndex := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, val := range scanner.Text() {
			switch val {
			case '.', '!', '?':
				if replace {
					fmt.Print(slang[slangIndex%len(slang)])

					replace = !replace
					slangIndex++

					continue
				}

				replace = !replace
			}

			fmt.Printf("%c", val)
		}
		fmt.Println()
	}

}

var slangIndex int
var slang = []string{
	", yeah!",
	", this is crazy, I tell ya.",
	", can U believe this?",
	", eh?",
	", aw yea.",
	", yo.",
	"? No way!",
	". Awesome!",
}

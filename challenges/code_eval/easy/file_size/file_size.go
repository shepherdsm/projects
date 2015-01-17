package main

import (
	"fmt"
	"os"
)

// File Size
// https://www.codeeval.com/open_challenges/26/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	stats, _ := file.Stat()

	fmt.Print(stats.Size())
}

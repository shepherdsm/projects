package main

import (
	"bufio"
	"fmt"
	"os"
)

// Vowel Count
// http://www.codeabbey.com/index/task_view/vowel-count
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var count int

		for _, val := range scanner.Text() {
			switch val {
			case 'a', 'o', 'u', 'i', 'e', 'y':
				count++
			}
		}

		fmt.Printf("%d ", count)
	}
}

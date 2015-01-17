package main

import (
	"fmt"
)

// Multiplication Tables
// https://www.codeeval.com/open_challenges/23
func main() {
	for i := 1; i <= 12; i++ {
		for k := 1; k <= 12; k++ {
			fmt.Printf("%4d", i*k)
		}
		fmt.Println()
	}
}

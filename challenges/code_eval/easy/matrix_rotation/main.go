package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Matrix Rotation
// https://www.codeeval.com/open_challenges/178/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix := strings.Split(scanner.Text(), " ")
		eles := len(matrix)
		row := int(math.Sqrt(float64(eles)))

		for i := row; i > 0; i-- {
			for k := eles - i; k >= 0; k -= row {
				fmt.Print(matrix[k], " ")
			}
		}

		fmt.Println()
	}
}

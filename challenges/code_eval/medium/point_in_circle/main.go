package main

import (
	"bufio"
	"fmt"
	"os"
)

// Point In Circle
// https://www.codeeval.com/open_challenges/98/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	format := "Center: (%f, %f); Radius: %f; Point: (%f, %f)"
	var (
		centerX float64
		centerY float64
		radius  float64
		pointX  float64
		pointY  float64
	)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), format, &centerX, &centerY, &radius, &pointX, &pointY)
		fmt.Println(centerX, centerY, radius, pointX, pointY)
	}
}

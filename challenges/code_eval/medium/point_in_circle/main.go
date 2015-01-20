package main

import (
	"bufio"
	"fmt"
	"math"
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
		fmt.Println(pointInCircle(centerX, centerY, radius, pointX, pointY))
	}
}

// Determines if a point is inside the circle by calculating the distance
// between the center and the point. If the center is less than the radius,
// we're in the circle.
func pointInCircle(cirX, cirY, rad, x, y float64) bool {
	d := math.Pow(x-cirX, 2) + math.Pow(y-cirY, 2)
	return d < math.Pow(rad, 2)
}

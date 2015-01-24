package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Fahrenheit to Celsius
// http://www.codeabbey.com/index/task_view/fahrenheit-celsius
func main() {
	for _, val := range os.Args[1:] {
		tmp, _ := strconv.Atoi(val)

		fmt.Printf("%d ", farToCelsius(tmp))
	}
}

func farToCelsius(fahr int) int {
	return round(float64(fahr-32) / 1.8)
}

func round(num float64) int {
	return int(math.Floor(num + 0.5))
}

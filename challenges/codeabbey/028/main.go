package main

import (
	"bufio"
	"fmt"
	"os"
)

// Body Mass Index
// http://www.codeabbey.com/index/task_view/body-mass-index
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var (
		weight float64
		height float64
	)

	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &weight, &height)
		fmt.Print(NewBMI(height, weight), " ")
	}
}

type BMI float64

func NewBMI(height, weight float64) BMI {
	return BMI(weight / (height * height))
}

func (b BMI) String() string {
	switch {
	case b < 18.5:
		return "under"
	case 18.5 <= b && b < 25.0:
		return "normal"
	case 25.0 <= b && b < 30.0:
		return "over"
	default:
		return "obese"
	}
}

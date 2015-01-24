package main

import (
	"fmt"
	"os"
	"strconv"
)

// Maximum of arry
// http://www.codeabbey.com/index/task_view/maximum-of-array
func main() {
	var largest int
	min := 0xFFFF

	for _, val := range os.Args[1:] {
		tmp, _ := strconv.Atoi(val)
		if tmp > largest {
			largest = tmp
		}
		if tmp < min {
			min = tmp
		}
	}

	fmt.Println(largest, min)
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

// Sum in Loop
// http://www.codeabbey.com/index/task_view/sum-in-loop
func main() {
	var sum int

	for _, val := range os.Args[1:] {
		tmp, _ := strconv.Atoi(val)
		sum += tmp
	}

	fmt.Println(sum)
}

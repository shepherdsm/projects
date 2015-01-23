package main

import (
	"fmt"
	"os"
	"strconv"
)

// Array Checksum
// http://www.codeabbey.com/index/task_view/array-checksum
func main() {
	var result int

	for _, val := range os.Args[1:] {
		num, _ := strconv.Atoi(val)

		result += num
		result *= 113
		result %= 10000007
	}

	fmt.Println(result)

}

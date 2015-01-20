package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Array Absurdity
// https://www.codeeval.com/open_challenges/41/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), ";")
		n, _ := strconv.Atoi(lines[0])

		// Sum up the array. Since it's 0..N-2, we can substract
		// out the sum of 1..N-2 to find out the duplicated number.
		var sum int
		for _, item := range strings.Split(lines[1], ",") {
			num, _ := strconv.Atoi(item)
			sum += num
		}

		fmt.Println(sum - natSum(n-2))
	}
}

func natSum(n int) int {
	return (n * (n + 1)) / 2
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Armstrong Numbers
// https://www.codeeval.com/open_challenges/82/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(isArmstrong(num))
	}
}

func isArmstrong(n int) string {
	if n < 10 {
		return "True"
	}

	if n == digitSum(n) {
		return "True"
	} else {
		return "False"
	}
}

func digitSum(n int) int {
	var sum float64
	digs := numDigits(n)

	for n > 0 {
		sum += math.Pow(float64(n%10), digs)
		n /= 10
	}

	return int(sum)
}

func numDigits(n int) float64 {
	return math.Floor(math.Log10(float64(n)) + 1)
}

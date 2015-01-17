package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Reverse And Add
// https://www.codeeval.com/open_challenges/45/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(revAndAdd(num))
	}
}

func isPalindrome(n int) bool {
	sNum := strconv.Itoa(n)
	l := len(sNum)

	for k := 0; k < l/2; k++ {
		if sNum[k] != sNum[l-k-1] {
			return false
		}
	}

	return true
}

func reverseNumber(n int) int {
	var revNum int

	for n > 0 {
		revNum *= 10
		revNum += n % 10
		n /= 10
	}

	return revNum
}

func revAndAdd(n int) (int, int) {
	var count int

	for !isPalindrome(n) {
		rev := reverseNumber(n)
		n += rev
		count++
	}

	return count, n
}

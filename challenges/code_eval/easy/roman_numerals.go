package main

import (
	"bufio"
	"fmt"
	"os"
)

// Roman Numerals
// https://www.codeeval.com/open_challenges/106/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var num int
		fmt.Sscanf(scanner.Text(), "%d", &num)
		fmt.Println(decToRoman(num))
	}
}

func decToRoman(n int) string {
	mapping := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	current := numbers[0]
	index := 1
	var res string

	for n > 0 {
		if current <= n {
			res += mapping[current]
			n -= current
		} else {
			current = numbers[index]
			index++
		}
	}

	return res
}

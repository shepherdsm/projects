package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Nice Angles
// https://www.codeeval.com/open_challenges/160/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.ParseFloat(scanner.Text(), 64)
		fmt.Println(convert(num))
	}
}

func convert(f float64) string {
	degrees := int(f)

	res := (f - float64(degrees)) * 60
	minutes := int(res)

	seconds := int((res - float64(minutes)) * 60)

	return fmt.Sprintf("%d.%02d'%02d\"", degrees, minutes, seconds)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Swap Elements
// https://www.codeeval.com/open_challenges/112/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " : ")
		arr := strings.Split(items[0], " ")

		for _, val := range strings.Split(items[1], ", ") {
			var first, second int
			fmt.Sscanf(val, "%d-%d", &first, &second)

			arr[first], arr[second] = arr[second], arr[first]
		}

		fmt.Println(strings.Join(arr, " "))
	}
}

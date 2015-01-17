package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type ItemInfo struct {
	Id    int
	Label string
}

type HeaderInfo struct {
	Header string
	Items  []ItemInfo
}

type MenuInfo struct {
	Menu HeaderInfo
}

// Json Menu IDs
// https://www.codeeval.com/open_challenges/102/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var menu MenuInfo
		blob := []byte(scanner.Text())
		json.Unmarshal(blob, &menu)

		var sum int
		for _, item := range menu.Menu.Items {
			if item.Label != "" {
				sum += item.Id
			}
		}

		fmt.Println(sum)
	}
}

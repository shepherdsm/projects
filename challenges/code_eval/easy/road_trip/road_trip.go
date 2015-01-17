package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Road Trip
// https://www.codeeval.com/open_challenges/124/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(cityRoute(strings.TrimRight(scanner.Text(), ";")))
	}
}

func cityRoute(s string) string {
	cities := strings.Split(s, ";")
	routes := make([]int, len(cities))

	for ind, city := range cities {
		items := strings.Split(city, ",")
		distance, _ := strconv.Atoi(items[1])
		routes[ind] = distance
	}

	sort.Ints(routes)
	prev := routes[0]
	results := make([]string, len(routes))
	results[0] = strconv.Itoa(prev)

	for ind, val := range routes[1:] {
		results[ind+1] = strconv.Itoa(val - prev)
		prev = val
	}

	return strings.Join(results, ",")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		fmt.Println(cityRoute(scanner.Text()))
	}
}

func cityRoute(s string) string {
	pat := regexp.MustCompile(",\\d+;")
	distances := pat.FindAllString(s, -1)
	routes := make([]int, len(distances))

	for ind, dist := range distances {
		l := len(dist)
		num, _ := strconv.Atoi(dist[1 : l-1])
		routes[ind] = num
	}

	sort.Ints(routes)

	prev := routes[0]
	res := make([]string, len(routes))
	res[0] = strconv.Itoa(prev)

	for ind, val := range routes[1:] {
		res[ind+1] = strconv.Itoa(val - prev)
		prev = val
	}

	return strings.Join(res, ",")
}

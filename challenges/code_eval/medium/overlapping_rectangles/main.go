package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Overlapping Rectangles
// https://www.codeeval.com/open_challenges/70/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		r1 := NewRectString(line[:4])
		r2 := NewRectString(line[4:])

		if r1.Overlaps(r2) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}

type Rect struct {
	ulx int
	uly int
	lrx int
	lry int
}

// Expects str to be in the format: [ulx uly lrx lry]
func NewRectString(str []string) *Rect {
	ulx := getCoord(str[0])
	uly := getCoord(str[1])
	lrx := getCoord(str[2])
	lry := getCoord(str[3])

	return &Rect{ulx, uly, lrx, lry}
}

func (r1 *Rect) Overlaps(r2 *Rect) bool {
	return !(r1.lrx < r2.ulx ||
		r1.lry > r2.uly ||
		r1.ulx > r2.lrx ||
		r1.uly < r2.lry)
}

func getCoord(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

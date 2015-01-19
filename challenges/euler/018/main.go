package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Maximum Path Sum II
// https://projecteuler.net/problem=18
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	triangle := NewTriangle()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		triangle.AddRowString(scanner.Text())
	}
	fmt.Println(triangle.MaxValue())
}

type Triangle struct {
	rows []row
}

type row []int

func NewTriangle() *Triangle {
	return &Triangle{make([]row, 0)}
}

func (t *Triangle) AddRowString(s string) {
	nums := strings.Split(s, " ")
	row := make(row, len(nums))

	for ind, val := range nums {
		tmp, _ := strconv.Atoi(val)
		row[ind] = tmp
	}

	t.rows = append(t.rows, row)
}

// Determine the maximum value by walking up the triangle from the bottom.
func (t *Triangle) MaxValue() int {
	// Start on the second to last row and walk backwards.
	// We pick the largest value from the previous row that's adjacent,
	// and add it to the current column.
	// Eventually we'll bubble up the largest value to t.rows[0][0]
	for k := len(t.rows) - 2; k >= 0; k-- {
		for col, val := range t.rows[k] {
			prev := t.rows[k+1]
			if left, right := prev[col], prev[col+1]; left > right {
				val += left
			} else {
				val += right
			}

			t.rows[k][col] = val
		}
	}

	return t.rows[0][0]
}

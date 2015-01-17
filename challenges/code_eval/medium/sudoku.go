package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SudokuGrid struct {
	side  int
	group int
	board []int
}

func (sg *SudokuGrid) IsValid() bool {
	for k := 0; k < sg.side; k++ {
		if !sg.IsValidRow(k) ||
			!sg.IsValidCol(k) ||
			!sg.IsValidGroup(k) {
			return false
		}
	}

	return true
}

func (sg *SudokuGrid) IsValidRow(rowIndex int) bool {
	offset := rowIndex * sg.side
	var sum int

	for k := 0; k < sg.side; k++ {
		sum += sg.board[k+offset]
	}

	return sum == natSum(sg.side)
}

func (sg *SudokuGrid) IsValidCol(colIndex int) bool {
	var sum int

	for k := 0; k < sg.side; k++ {
		sum += sg.board[k*sg.side+colIndex]
	}

	return sum == natSum(sg.side)
}

func (sg *SudokuGrid) IsValidGroup(groupIndex int) bool {
	var sum int
	colOffset := (groupIndex % sg.group) * sg.group

	// Relies on integer division truncating floating point from division.
	rowOffset := (groupIndex / sg.group) * sg.group

	for r := rowOffset; r < rowOffset+sg.group; r++ {
		for c := colOffset; c < colOffset+sg.group; c++ {
			sum += sg.board[r*sg.side+c]
		}
	}

	return sum == natSum(sg.side)
}

// Sudoku
// https://www.codeeval.com/open_challenges/78/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sg := populateGrid(scanner.Text())

		if sg.IsValid() {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}

func populateGrid(s string) *SudokuGrid {
	sideToGroup := map[int]int{
		4: 2,
		9: 3,
	}
	sg := &SudokuGrid{}

	items := strings.Split(s, ";")
	side, _ := strconv.Atoi(items[0])

	sg.side = side
	sg.group = sideToGroup[side]
	sg.board = make([]int, side*side)

	for ind, val := range strings.Split(items[1], ",") {
		tmp, _ := strconv.Atoi(val)
		sg.board[ind] = tmp
	}

	return sg
}

func natSum(n int) int {
	return (n * (n + 1)) / 2
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	rows int
	cols int
	vals []int
}

func (b *Board) SumRow(rowIndex int) int {
	offset := b.rows * rowIndex
	var sum int

	for k := 0; k < b.cols; k++ {
		sum += b.vals[offset+k]
	}

	return sum
}

func (b *Board) SumCol(colIndex int) int {
	var sum int

	for k := 0; k < b.rows; k++ {
		sum += b.vals[b.rows*k+colIndex]
	}

	return sum
}

func (b *Board) SetCol(colIndex int, val int) {
	for k := 0; k < b.rows; k++ {
		b.vals[b.rows*k+colIndex] = val
	}
}

func (b *Board) SetRow(rowIndex int, val int) {
	offset := b.rows * rowIndex

	for k := 0; k < b.cols; k++ {
		b.vals[offset+k] = val
	}
}

// Query Board
// https://www.codeeval.com/open_challenges/87/
func main() {
	board := &Board{256, 256, make([]int, 256*256)}
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runCommand(scanner.Text(), board)
	}
}

func oneArg(s string) int {
	num1, _ := strconv.Atoi(s)

	return num1
}

func twoArgs(args []string) (int, int) {
	return oneArg(args[1]), oneArg(args[2])
}

func runCommand(s string, b *Board) {
	items := strings.Split(s, " ")

	switch strings.ToUpper(items[0]) {
	case "SETCOL":
		b.SetCol(twoArgs(items))
	case "SETROW":
		b.SetRow(twoArgs(items))
	case "QUERYCOL":
		fmt.Println(b.SumCol(oneArg(items[1])))
	case "QUERYROW":
		fmt.Println(b.SumRow(oneArg(items[1])))
	}
}

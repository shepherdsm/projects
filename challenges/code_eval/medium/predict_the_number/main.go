package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Predict The Number
// https://www.codeeval.com/open_challenges/125/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	seq := newSequence()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		fmt.Println(seq.get(num))
	}
}

type Sequence []uint8

func newSequence() *Sequence {
	return &Sequence{0}
}

func (seq *Sequence) get(n int) uint8 {
	if len(*seq) < n {
		seq = seq.expand(n)
	}

	return []uint8(*seq)[n]
}

func (seq Sequence) expand(n int) *Sequence {
	newCap := int(math.Pow(2, math.Ceil(math.Log2(float64(n)))))
	newSlice := make(Sequence, len(seq), newCap)
	copy(newSlice, seq)

	for len(newSlice) < newCap {
		for _, val := range newSlice {
			newSlice = append(newSlice, changeNum(val))
		}
	}

	return &newSlice
}

func changeNum(n uint8) uint8 {
	switch n {
	case 0:
		return 1
	case 1:
		return 2
	default:
		return 0
	}
}

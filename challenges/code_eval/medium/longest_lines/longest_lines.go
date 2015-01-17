package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type LineHeap []string

func (h LineHeap) Len() int           { return len(h) }
func (h LineHeap) Less(i, j int) bool { return len(h[i]) > len(h[j]) }
func (h LineHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *LineHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *LineHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}

// Longest Lines
// https://www.codeeval.com/open_challenges/2/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())

	h := &LineHeap{}
	heap.Init(h)

	for scanner.Scan() {
		heap.Push(h, scanner.Text())
	}

	for k := 0; k < num; k++ {
		fmt.Println(heap.Pop(h))
	}
}

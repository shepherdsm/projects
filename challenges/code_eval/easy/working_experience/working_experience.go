package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type WorkSpan struct {
	start time.Time
	end   time.Time
}

type WorkList []WorkSpan

func (l WorkList) Len() int {
	return len(l)
}

func (l WorkList) Less(i, j int) bool {
	span1 := l[i]
	span2 := l[j]

	if span1.start.Equal(span2.start) {
		return span1.end.Before(span2.end)
	}

	return span1.start.Before(span2.start)
}

func (l WorkList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Working Experience
// https://www.codeeval.com/open_challenges/139/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(getActualWork(scanner.Text()))
	}
}

func combineAdjacent(work WorkList) WorkList {
	var newList WorkList

	span := work[0]
	for _, val := range work[1:] {
		if val.start.After(span.start) && val.end.Before(span.end) ||
			val.start.Equal(span.start) && val.end.Equal(span.end) {
			continue // Embedded time range should be skipped.
		} else if val.start.Before(span.end) || span.end.Equal(val.start) {
			span.end = val.end
		} else {
			newList = append(newList, span)
			span = val
		}
	}

	newList = append(newList, span)

	return newList
}

func getActualWork(s string) int {
	dates := strings.Split(s, "; ")
	dateStr := "Jan 2006"
	workExp := make(WorkList, len(dates))

	for ind, val := range dates {
		tmp := strings.Split(val, "-")

		s, _ := time.Parse(dateStr, tmp[0])
		e, _ := time.Parse(dateStr, tmp[1])

		workExp[ind] = WorkSpan{s, e}
	}

	sort.Sort(workExp)

	var totalMonths int
	for _, val := range combineAdjacent(workExp) {
		count := 1
		for val.start.Before(val.end) {
			val.start = val.start.AddDate(0, 1, 0)
			count++
		}

		totalMonths += count
	}

	return totalMonths / 12
}

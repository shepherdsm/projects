package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Delta Time
// https://www.codeeval.com/open_challenges/166/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), " ")
		fmt.Println(getDelta(items[0], items[1]))
	}
}

func getDelta(t1, t2 string) string {
	timeStr := "15:04:05"
	time1, _ := time.Parse(timeStr, t1)
	time2, _ := time.Parse(timeStr, t2)
	var dur time.Duration

	if time1.After(time2) {
		dur = time1.Sub(time2)
	} else {
		dur = time2.Sub(time1)
	}

	hours := int(dur.Hours())
	minutes := int(dur.Minutes()) - (hours * 60)
	seconds := (int(dur.Seconds()) - (minutes * 60)) % 100
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

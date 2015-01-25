package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Triangle Area
// http://www.codeabbey.com/index/task_view/triangle-area
func main() {
	var (
		x1 float64
		y1 float64
		x2 float64
		y2 float64
		x3 float64
		y3 float64
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &x1, &y1, &x2, &y2, &x3, &y3)

		p1, p2, p3 := &Point{x1, y1}, &Point{x2, y2}, &Point{x3, y3}
		t := &Triangle{p1.Distance(p2), p2.Distance(p3), p3.Distance(p1)}

		fmt.Printf("%.1f ", t.Area())
	}
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t *Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

type Point struct {
	x float64
	y float64
}

func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

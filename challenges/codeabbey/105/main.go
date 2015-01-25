package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Convex Polygon Area
// http://www.codeabbey.com/index/task_view/convex-polygon-area
func main() {
	var (
		x float64
		y float64
	)

	poly := NewPolygon()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Sscan(scanner.Text(), &x, &y)
		poly.AddPoint(&Point{x, y})
	}

	fmt.Printf("%.1f\n", poly.Area())
}

type Polygon struct {
	points []*Point
}

func NewPolygon() *Polygon {
	return &Polygon{make([]*Point, 0)}
}

func (poly *Polygon) AddPoint(p *Point) {
	poly.points = append(poly.points, p)
}

// Slices the polygon into a series of triangles based on a vertex and sums their area.
func (poly *Polygon) Area() float64 {
	vertex := poly.points[0]
	var area float64

	for i := 1; i < len(poly.points)-1; i++ {
		p2, p3 := poly.points[i], poly.points[i+1]
		t := &Triangle{vertex.Distance(p2), p2.Distance(p3), p3.Distance(vertex)}
		area += t.Area()
	}

	return area
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

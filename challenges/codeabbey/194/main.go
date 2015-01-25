package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
)

var gStartPost *Post

// Convex Hull and Farmers
// http://www.codeabbey.com/index/task_view/convex-hull-and-farmers
func main() {
	posts, _ := getConvexHull()

	for _, val := range posts {
		fmt.Printf("%d ", val.ind)
	}
}

// Loads all of the posts into an array and sets the global gStartPost
func getPosts() Posts {
	var totalPosts int
	fmt.Fscan(os.Stdin, &totalPosts)
	posts := make(Posts, 0, totalPosts)

	var (
		x int
		y int
	)

	fmt.Fscan(os.Stdin, &x, &y)
	gStartPost = &Post{&Point{x, y}, 0}
	posts = append(posts, gStartPost)

	for i := 1; i < totalPosts; i++ {
		fmt.Fscan(os.Stdin, &x, &y)
		curPost := &Post{&Point{x, y}, i}

		posts = append(posts, curPost)

		if curPost.greater(gStartPost) {
			gStartPost = curPost
		}
	}

	return posts
}

// Runs the Graham Scan algorithm to find the Convex Hull for the set of points.
func getConvexHull() (Posts, error) {
	posts := getPosts()
	sort.Sort(posts)
	totalPoints := len(posts)

	posts[0] = posts[totalPoints-1]

	numPoints := 1

	for i := 2; i < totalPoints; i++ {
		for isLeftTurn(posts[numPoints-1], posts[numPoints], posts[i]) > 0 {
			if numPoints > 1 {
				numPoints -= 1
			} else if i == totalPoints {
				break
			} else {
				i++
			}
		}

		numPoints++
		posts[numPoints], posts[i] = posts[i], posts[numPoints]
	}

	for ind, val := range posts[1:] {
		if posts[0].ind == val.ind {
			posts[0] = gStartPost
			return posts[:ind+2], nil
		}
	}

	return nil, errors.New("No convex hull found.")
}

type Posts []*Post

type Post struct {
	p   *Point
	ind int
}

type Point struct {
	x int
	y int
}

// Determines if moving from p1 -> p2 -> p3 consitutes a left turn or right turn.
func isLeftTurn(p1, p2, p3 *Post) int {
	turn := (p2.p.x-p1.p.x)*(p3.p.y-p1.p.y) - (p2.p.y-p1.p.y)*(p3.p.x-p1.p.x)

	return turn
}

// Used to determine the starting point.
func (p1 *Post) greater(p2 *Post) bool {
	return (p1.p.y > p2.p.y) ||
		(p1.p.y == p2.p.y && p1.p.x < p2.p.x)
}

// Returns if p1 has a greater angle compared to the startPost than p2.
func (p1 *Post) greaterAngle(p2 *Post) bool {
	angle1 := p1.angle()
	angle2 := p2.angle()

	return angle1 > angle2
}

// Determines the angle between the x-axis and the line from gStartPoint -> p1
func (p1 *Post) angle() float64 {
	deltaY := float64(p1.p.y - gStartPost.p.y)
	deltaX := float64(p1.p.x - gStartPost.p.x)

	return math.Atan2(deltaY, deltaX)
}

// Define the interface for sort.Sort
func (p Posts) Len() int {
	return len(p)
}

func (p Posts) Less(i, j int) bool {
	return p[i].greaterAngle(p[j])
}

func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

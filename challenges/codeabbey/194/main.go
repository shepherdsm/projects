package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var gStartPost *Post

// Convex Hull and Farmers
// http://www.codeabbey.com/index/task_view/convex-hull-and-farmers
func main() {
	posts := getConvexHull()

	posts.printPosts()
}

// Ugliest printing method ever.
func (p Posts) printPosts() {
	hull := make(Posts, 0, len(p))
	end := p[0]
	hull = append(hull, gStartPost)

	// Pull out the points that define the convex hull.
	for _, val := range p[1:] {
		if end.ind == val.ind {
			break
		}

		hull = append(hull, val)
	}

	hull = append(hull, end)

	var (
		start   int
		largest int
	)

	// Find the largest y value to be our start.
	for ind, val := range hull {
		if val.p.y > largest {
			largest = val.p.y
			start = ind
		}
	}

	index := start

	// We'll print in clockwise order, but because of how we choose
	// our starting point for the Graham Scan algorithm, we may be in
	// the middle of the array.
	for {
		fmt.Printf("%d ", hull[index].ind)
		index--

		if index == -1 {
			index = len(hull) - 1
		}
		if index == start {
			break
		}
	}
}

// Loads all of the posts into an array and sets the global gStartPost
func getPosts() Posts {
	scanner := bufio.NewScanner(os.Stdin)

	var totalPosts int

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "&d", &totalPosts)
	posts := make(Posts, 0, totalPosts)

	var (
		currentPost int
		x           int
		y           int
	)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
		curPost := &Post{&Point{x, y}, currentPost}

		posts = append(posts, curPost)

		if gStartPost == nil {
			gStartPost = curPost
		}

		if curPost.less(gStartPost) {
			gStartPost = curPost
		}

		currentPost++
	}

	return posts
}

// Runs the Graham Scan algorithm to find the Convex Hull for the set of points.
func getConvexHull() Posts {
	posts := getPosts()
	sort.Sort(posts)
	totalPoints := len(posts)

	posts[0] = posts[totalPoints-1]

	numPoints := 1

	for i := 2; i < totalPoints; i++ {
		for isLeftTurn(posts[numPoints-1], posts[numPoints], posts[i]) <= 0 {
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

	return posts
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
func (p1 *Post) less(p2 *Post) bool {
	return (p1.p.y < p2.p.y) ||
		(p1.p.y == p2.p.y && p1.p.x < p2.p.x)
}

// Returns if p1 has a smaller angle compared to the startPost than p2.
func (p1 *Post) lessAngle(p2 *Post) bool {
	angle1 := p1.angle()
	angle2 := p2.angle()

	return angle1 < angle2
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
	return p[i].lessAngle(p[j])
}

func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

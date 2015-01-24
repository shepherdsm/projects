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
	posts := getPosts()
	sort.Sort(posts)
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

type Posts []*Post

type Post struct {
	p   *Point
	ind int
}

type Point struct {
	x int
	y int
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

// Determines the angle between the current post and the global startPost.
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

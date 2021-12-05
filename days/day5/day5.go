package day5

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
)


func Day5(inputFile string, part int) {
	if part == 0 {
		day5(inputFile, false)
	} else {
		day5(inputFile, true)
	}
}

type Point struct {
	x	int
	y   int
}

func day5(inputFile string, traverseAll bool) {
	ws := util.LineScanner(inputFile)
	read, hasNext := util.Read(ws)

	// Assuming 1000x1000 grid
	coverageMap := make([][]int, 1000)
	for i := range coverageMap {
		coverageMap[i] = make([]int, 1000)
	}

	var (
		a       Point
		b       Point
		c       Point
		isctPts = 0
	)
	// Read in each segment
	for hasNext {
		a, b = toEndpoints(read) // promises a.x <= b.x
		c = Point{x: -1, y: -1}
		// traverse each segment, or horizontal+vertical segments only
		if traverseAll || a.x == b.x || a.y == b.y {
			reachedEnd := false
			for !reachedEnd {
				// First traversal starts at point a
				if c.x == -1 {
					c.x = a.x
					c.y = a.y
				} else {
					// Otherwise move point c +1 in each direction
					if a.x < b.x {
						c.x++
					}
					if a.y < b.y {
						c.y++
					} else if a.y > b.y {
						c.y--
					}
				}

				// Check if we've traversed to the endpoint
				if c.x == b.x && c.y == b.y {
					reachedEnd = true
				}

				coverageMap[c.x][c.y]++
				// Only bump if we just found the first intersect
				if coverageMap[c.x][c.y] == 2 {
					isctPts++
				}
			}

		}

		read, hasNext = util.Read(ws)
	}

	fmt.Printf("Intersecting points: %d\n", isctPts)
}

func toEndpoints(read string) (Point, Point) {
	coords := strings.Split(read, " -> ")
	pt1 := strings.Split(coords[0], ",")
	pt2 := strings.Split(coords[1], ",")

	p1 := Point{
		x: util.ToInt(pt1[0]),
		y: util.ToInt(pt1[1]),
	}

	p2 := Point{
		x: util.ToInt(pt2[0]),
		y: util.ToInt(pt2[1]),
	}

	if p1.x <= p2.x {
		return p1, p2
	} else {
		return p2, p1
	}
}
package days

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
)

type Line struct {
	x1 	int
	y1   int
	x2   int
	y2   int
}

func Day5a(inputFile string) {
	ws := util.LineScanner(inputFile)
	read, ok := util.Read(ws)

	isctPts := 0
	var line Line
	var horizOrVertical bool
	coverageMap := make([][]int, 1000)
	for i := range coverageMap {
		coverageMap[i] = make([]int, 1000)
	}
	for ok {
		line, horizOrVertical = toLine(read)
		if horizOrVertical {
			isctPts += mark(coverageMap, line)
		}

		read, ok = util.Read(ws)
	}

	fmt.Printf("Points which intersect at least 2 lines: %d", isctPts)
}

func mark(cmap [][]int, line Line) int {
	newIntersects := 0
	// vertical
	if line.x1 == line.x2 {
		for y := line.y1; y <= line.y2; y++ {
			cmap[line.x1][y]++
			if cmap[line.x1][y] == 2 {
				newIntersects++
			}
		}
	}

	// horizontal
	if line.y1 == line.y2 {
		for x := line.x1; x <= line.x2; x++ {
			cmap[x][line.y1]++
			if cmap[x][line.y1] == 2 {
				newIntersects++
			}
		}
	}

	return newIntersects
}

func toLine(read string) (Line, bool) {

	coords := strings.Split(read, " -> ")
	pt1 := strings.Split(coords[0], ",")
	pt2 := strings.Split(coords[1], ",")

	X1 := util.ToInt(pt1[0])
	Y1 := util.ToInt(pt1[1])
	X2 := util.ToInt(pt2[0])
	Y2 := util.ToInt(pt2[1])

	var (
		x1	int
		x2	int
		y1	int
		y2	int
	)
	if X1 == X2 {
		x1 = X1
		x2 = X2
		if Y1 <= Y2 {
			y1 = Y1
			y2 = Y2
		} else {
			y1 = Y2
			y2 = Y1
		}

		return Line{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}, true
	}

	if Y1 == Y2 {
		y1 = Y1
		y2 = Y2
		if X1 <= X2 {
			x1 = X1
			x2 = X2
		} else {
			x1 = X2
			x2 = X1
		}

		return Line{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}, true
	}

	// Skip line if it's not horizontal or vertical
	return Line{}, false
}
package days

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Day13(inputFile string, part int) {
	ls := util.LineScanner(inputFile)
	if part == 0 {
		day13a(ls)
	} else {
		day13b(ls)
	}
}

func day13a(ls *bufio.Scanner) {
	pts, folds, sizeX, sizeY := readInput(ls)

	singleFold := folds[0]
	sheet := makeSheet(sizeX, sizeY, []Fold{singleFold})

	totalMarks := 0
	for _, pt := range pts {
		x, y := singleFold.apply(pt.x, pt.y)
		if sheet[x][y] == 0 {
			sheet[x][y] = 1
			totalMarks++
		}
	}

	fmt.Printf("Total marks: %d\n", totalMarks)
}

func day13b(ls *bufio.Scanner) {
	pts, folds, sizeX, sizeY := readInput(ls)
	sheet := makeSheet(sizeX, sizeY, folds)

	for _, pt := range pts {
		x, y := pt.x, pt.y
		for _, fold := range folds {
			x, y = fold.apply(x,y)
		}
		if sheet[x][y] == 0 {
			sheet[x][y] = 1
		}
	}

	for j := 0; j < len(sheet[0]); j++ {
		for i := 0; i < len(sheet); i++ {
			if sheet[i][j] == 1 {
				fmt.Printf(" # ")
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Println("")
	}
}

type Pt struct {
	x 	int
	y 	int
}

type Fold struct {
	x 	int
	y 	int
}

func (f Fold) apply(x int, y int) (int, int) {
	return util.Abs(f.x - util.Abs(f.x-x)), util.Abs(f.y - util.Abs(f.y-y))
}

func makeSheet(sizeX int, sizeY int, folds []Fold) [][]int {
	for _, f := range folds {
		if f.x != 0 {
			sizeX = util.Max(f.x, sizeX-f.x-1)
		} else {
			sizeY = util.Max(f.y, sizeY-f.y-1)
		}
	}

	sheet := make([][]int, sizeX)
	for idx := range sheet {
		sheet[idx] = make([]int, sizeY)
	}

	return sheet
}

func readInput(ls *bufio.Scanner) ([]Pt, []Fold, int, int) {
	line, ok := util.Read(ls)

	var (
		x     int
		y     int
		maxX = 0
		maxY  = 0
		pts   = make([]Pt, 0)
	)

	for ok && line != "" {
		coords := strings.Split(line, ",")
		x, _ = strconv.Atoi(coords[0])
		y, _ = strconv.Atoi(coords[1])
		pts = append(pts, Pt{x, y})

		maxX = util.Max(maxX, x)
		maxY = util.Max(maxY, y)

		line, ok = util.Read(ls)
	}

	line, ok = util.Read(ls)
	folds := make([]Fold, 0)
	for ok {
		x, y = 0, 0
		if strings.Contains(line, "fold along x=") {
			x, _ = strconv.Atoi(strings.Split(line, "=")[1])
		} else {
			y, _ = strconv.Atoi(strings.Split(line, "=")[1])
		}

		folds = append(folds, Fold{x, y})
		line, ok = util.Read(ls)
	}

	return pts, folds, maxX+1, maxY+1
}

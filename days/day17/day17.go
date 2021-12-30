package day17

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func Day17(inputFile string, part int) {
	ls := util.LineScanner(inputFile)
	area := area(ls)
	area.findTrajectories()
	maxHeight := -1
	distinct := 0
	for _, v := range area.tops {
		for _, vv := range v {
			if vv > maxHeight {
				maxHeight = vv
			}
			distinct++
		}
	}

	if part == 0 {
		fmt.Printf("Max height: %d\n", maxHeight)
	} else {
		fmt.Printf("Distinct: %d \n", distinct)
	}
}

// Constraints: For a given time t and position (xt, yt):
//
// X VELOCITY: constrained with min and max drag
// (xt - t(t-1)/2)/t <= v.x <= (xt + t(t-1)/2)/t
//
// Y VELOCITY:
// yt = t*v.y - t(t-1)/2
// =>    v.y = (yt + t(t-1)/2)/t
//
// MAX Y HEIGHT:
// y't == 0 => v.y - t + 0.5 == 0 > t == v.y + 0.5 ~ v.y
// =>    occurs at step t = v.y
//
// MAX T:
// t <= 2*area.maxX
func (a *Area) findTrajectories() {
	for t := 1; t <= 2*a.maxX; t++ {
		// For every possible endpoint (xt, yt), find a starting velocity that works
		for xt := a.minX; xt <= a.maxX; xt++ {
			for yt := a.minY;  yt <= a.maxY; yt++ {
				yvel := (yt + t*(t-1)/2)/t
				if (yvel*t) == yt + t*(t-1)/2 { // Found valid y-velocity
					for drag := -t*(t-1)/2; drag <= t*(t-1)/2; drag++ {
						xvel := (xt + drag)/t
						if (xvel*t) == xt + drag { // Found possibly valid x-velocity
							xpos := simulateX(xvel, t)
							if xpos == xt { // verify
								if a.tops[xvel] == nil {
									a.tops[xvel] = make(map[int]int)
								}
								a.tops[xvel][yvel] = yvel*yvel - yvel*(yvel-1)/2
							}
						}
					}
				}
			}
		}
	}
}

func simulateX(xvel int, t int) int {
	xpos := 0
	for i := 1; i <= t; i++ {
		xpos += xvel
		if xvel > 0 {
			xvel--
		} else if xvel < 0 {
			xvel++
		}
	}
	return xpos
}

// PARSE INPUT
type Area struct {
	minX	int
	maxX	int
	minY	int
	maxY	int
	tops	map[int]map[int]int
}

func area(ls *bufio.Scanner) Area {
	line, _ := util.Read(ls)
	re := regexp.MustCompile("(-)*[0-9]+")
	limits := re.FindAllString(line, -1)
	minx, _ := strconv.Atoi(limits[0])
	maxx, _ := strconv.Atoi(limits[1])
	y1, _ := strconv.Atoi(limits[2])
	y2, _ := strconv.Atoi(limits[3])

	var miny, maxy int
	if y1 <= y2 {
		miny = y1
		maxy = y2
	} else {
		miny = y2
		maxy = y1
	}
	return Area{minx, maxx, miny, maxy, make(map[int]map[int]int)}
}

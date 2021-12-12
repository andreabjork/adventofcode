package days

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)


func Day9(inputFile string, part int) {
	if part == 0 {
		day9a(inputFile)
	} else {
		day9b(inputFile)
	}
}

type Basin struct {
	coords map[int]map[int]bool
	size int
}

type Coord struct {
	x 	int
	y 	int
}

func day9a(inputFile string) {
	_, _, risk := mapCave(inputFile)
	fmt.Printf("Risk: %d\n", risk)
}

func day9b(inputFile string) {
	cave, lmin, _ := mapCave(inputFile)

	emptyMap := make(map[int]map[int]bool)
	// explore out from every local minima
	biggestBasins := make([]*Basin, 0)
	for _, m := range lmin {
		basin := &Basin{emptyMap, 0}
		basin = explore(cave, basin, m)
		biggestBasins = evaluate(basin, biggestBasins)
	}

	multiple := 1
	for _, b := range biggestBasins {
		multiple *= (*b).size
	}

	fmt.Printf("3 biggest: %d\n", multiple)
}

// Evaluate whether basin b is among the three biggest
// basins found so far.
func evaluate(b *Basin, biggest []*Basin) []*Basin {
	if len(biggest) < 3 {
		biggest = append(biggest, b)
		return biggest
	}

	// Okay we could keep track of the minimum basin at all times...
	min := (*biggest[0]).size
	minIdx := 0
	for idx, big := range biggest {
		if (*big).size < min {
			minIdx = idx
			min = (*biggest[minIdx]).size
		}
	}

	if (*b).size > (*biggest[minIdx]).size {
		biggest[minIdx] = b
	}

	return biggest
}

// Explores from a point p in basin b.
// Explore visits all adjacent points.
// Explore returns if it's at a point with height 9,
// a point already in the basin or a point outside of the cave.
func explore(cave [][]int, b *Basin, p *Coord) *Basin {
	// stop if point is outside of the cave
	if (*p).x < 0 || (*p).x >= len(cave) || (*p).y < 0 || (*p).y >= len(cave[(*p).x]) {
		return b
	}

	// stop if we reach a 9
	if cave[(*p).x][(*p).y] == 9 {
		return b
	}

	// stop if point is already in basin
	if b.coords[(*p).x][(*p).y] {
		return b
	}

	// otherwise addCave this point to the basin and continue exploring
	if (*b).coords[(*p).x] == nil {
		(*b).coords[(*p).x] = make(map[int]bool)
	}
	(*b).coords[(*p).x][(*p).y] = true
	(*b).size++

	b = explore(cave, b, &Coord{p.x+1, p.y})
	b = explore(cave, b, &Coord{p.x-1, p.y})
	b = explore(cave, b, &Coord{p.x, p.y+1})
	b = explore(cave, b, &Coord{p.x, p.y-1})

	return b
}

// Bruteforce mapping of the cave; a point is a local
// minima if all its adjacent setNeighbours (up, down, right, left)
// are higher than it is.
func mapCave(inputFile string) ([][]int, []*Coord, int) {
	ws := util.LineScanner(inputFile)
	line, ok := util.Read(ws)

	var (
		runes	[]rune
		cave	= make([][]int, 0)
	)
	// read inputs
	for i := 0; ok; i++ {
		runes = []rune(line)
		cave = append(cave, make([]int, 0))
		for _, r := range runes {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				fmt.Printf("Something went wrong reading input.")
			}
			cave[i] = append(cave[i], num)
		}

		line, ok = util.Read(ws)
	}

	lMin := make([]*Coord, 0)
	sum := 0
	for i := 0; i < len(cave); i++ {
		for j := 0; j < len(cave[i]); j++ {
			min := true

			if i-1 >= 0 && cave[i-1][j] <= cave[i][j] {
				min = false
			}

			if i+1 < len(cave) && cave[i+1][j] <= cave[i][j] {
				min = false
			}

			if j-1 >= 0 && cave[i][j-1] <= cave[i][j] {
				min = false
			}

			if j+1 < len(cave[i]) && cave[i][j+1] <= cave[i][j] {
				min = false
			}

			if min {
				sum += cave[i][j] + 1
				if i < 2 {
				}
				lMin = append(lMin, &Coord{i, j})
			}
		}
	}

	return cave, lMin, sum
}

package day12

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
	"strings"
)

func Day12(inputFile string, part int) {
	ls := util.LineScanner(inputFile)
	if part == 0 {
		day12a(ls)
	} else {
		day12b(ls)
	}
}

func day12a(ls *bufio.Scanner) {
	cs := mapCaves(ls, 0)
	cs.exploreFrom((*cs).paths[0])

	fmt.Printf("Completed Paths: %d\n", (*cs).completedPaths)
}

func day12b(ls *bufio.Scanner) {
	cs := mapCaves(ls, 1)
	cs.exploreFrom((*cs).paths[0])

	fmt.Printf("Completed Paths: %d\n", (*cs).completedPaths)
}

// ===================
// MAP INPUTS TO GRAPH
// ===================
func mapCaves(ls *bufio.Scanner, bonusVisits int) *CaveSystem {
	cs := &CaveSystem{make(map[string]*Cave), make([]*Path, 0), 0}
	line, ok := util.Read(ls)
	for ok {
		input := strings.Split(line, "-")
		cs.addCave(input[0])
		cs.addCave(input[1])
		cs.setNeighbours(input[0], input[1])

		line, ok = util.Read(ls)
	}
	// Add the first path, including only the "start" cave:
	start := &Path{[]*Cave{(*cs).caves["start"]}, make(map[*Cave]int), bonusVisits}
	(*cs).paths = append((*cs).paths, start)
	return cs
}


// ==========
// CAVESYSTEM
// ==========
type CaveSystem struct {
	caves         map[string]*Cave
	paths         []*Path
	completedPaths int
}

func (cs *CaveSystem) addCave(name string) {
	var caveType CaveType
	if _, found := (*cs).caves[name]; !found {
		caveType = fromString(name)
		(*cs).caves[name] = &Cave{name, make([]*Cave, 0), caveType}
	}
}

func (cs *CaveSystem) setNeighbours(a string, b string) {
	caveA := (*cs).caves[a]
	caveB := (*cs).caves[b]

	(*caveA).neighbours = append((*caveA).neighbours, caveB)
	(*caveB).neighbours = append((*caveB).neighbours, caveA)
}

func (cs *CaveSystem) exploreFrom(p *Path) {
	// Splits path p to p1, p2, ..., pN where
	// p1 = p, {p.curr -> p.curr.neighbour(1)
	// p2 = p, {p.curr -> p.curr.neighbour(1)
	// ...
	// for each of p.curr's N neighbours.

	// If we're already at the end, no more exploring
	if p.curr().cType == End {
		return
	}

	var path *Path
	for idx, n := range (*p.curr()).neighbours {
		if (*n).cType == Big || (*n).cType == End || ((*n).cType == Small && ((*p).visited[n] < 1 || (*p).bonusVisits > 0 )) {
			if (*n).cType == End {
				(*cs).completedPaths++
			}

			if idx == len(p.curr().neighbours)-1 {
				// Extend this path
				path = p
				p.visit(n)
			} else {
				// Make new path
				path = (*p).copy()
				(*cs).paths = append((*cs).paths, path)
				path.visit(n)
			}
			cs.exploreFrom(path)
		}
	}
}

// =====
// CAVES
// =====
type CaveType int64
const (
	Start CaveType	= iota
	End
	Big
	Small
)

func fromString(s string) CaveType {
	if s == "start" {
		return Start
	} else if s == "end" {
		return End
	} else if strings.ToUpper(s) == s {
		return Big
	} else {
		return Small
	}
}

type Cave struct {
	name 		string
	neighbours 	[]*Cave
	cType 		CaveType
}

// =====
// PATHS
// =====
type Path struct {
	path			[]*Cave
	visited 		map[*Cave]int
	bonusVisits		int
}

func (p *Path) copy() *Path {
	newPath := Path{make([]*Cave, len((*p).path)), make(map[*Cave]int), (*p).bonusVisits}

	for i, c := range (*p).path {
		newPath.path[i] = c
	}

	for k, v := range (*p).visited {
		newPath.visited[k] = v
	}

	return &newPath
}

func (p *Path) curr() *Cave {
	return (*p).path[len((*p).path)-1]
}

func (p *Path) visit(c *Cave) {
	(*p).visited[c]++
	// If we visited a small node, note if a bonus visit was used
	if (*c).cType == Small && (*p).visited[c] > 1 {
		(*p).bonusVisits--
	}
	(*p).path = append((*p).path, c)
}
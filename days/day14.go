package days

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
	"strings"
)

func Day14(inputFile string, part int) {
	ls := util.LineScanner(inputFile)
	if part == 0 {
		makePolymer(ls, 10)
	} else {
		makePolymer(ls, 40)
	}
}

func makePolymer(ls *bufio.Scanner, steps int) {
	line, ok := util.Read(ls)
	// Create polymer
	template := []rune(line)
	p := Polymer{steps,  make(map[rune]map[rune]rune), make(map[rune]map[rune][]map[rune]int)}

	// Capture rules
	var (
		left	[]rune
		right 	[]rune
		parts	[]string
	)
	line, ok = util.Read(ls)
	line, ok = util.Read(ls)
	for ok {
		parts = strings.Split(line, " -> ")
		left = []rune(parts[0])
		right = []rune(parts[1])

		if p.rules[left[0]] == nil {
			p.rules[left[0]]  = make(map[rune]rune)
		}
		p.rules[left[0]][left[1]] = right[0]
		line, ok = util.Read(ls)
	}

	// Expand string
	c := make(map[rune]int)
	for i, j := 0, 1; j < len(template); i,j = i+1, j+1 {
		c = add(c, p.count(template[i], template[j], p.steps))
		if j < len(template) -1 {
			// remove duplicate count
			c = add(c, map[rune]int{template[j]: -1})
		}
	}

	min, max := int(^uint(0) >> 1), -1
	for _, v := range c {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	fmt.Println("Most common - least common: ", max - min)
}

type Polymer struct {
	steps     int
	rules     map[rune]map[rune]rune
	expansion map[rune]map[rune][]map[rune]int // expansion[x][y][3][z] = the number of z's added from XY after 3 steps
}

func (p *Polymer) count(a rune, b rune, step int) map[rune]int {
	// With xy -> z we get
	// expansion[x][y][3] = expansion[x][z][2]+expansion[z][y][2]
	c := p.rules[a][b]
	if step == 1 {
		cc := make(map[rune]int)
		cc[a]++
		cc[b]++
		cc[c]++
		return cc
	}

	// if XY expansion was never counted, create the counters
	if p.expansion[a] == nil {
		p.expansion[a] = make(map[rune][]map[rune]int)
	}
	if p.expansion[a][b] == nil {
		p.expansion[a][b] = make([]map[rune]int, p.steps+1)
	}

	// if XY expansion is not recorded at step 'step', count it now
	if p.expansion[a][b][step] == nil {
		p.expansion[a][b][step] = add(p.count(a, c, step-1), p.count(c, b, step-1))
		p.expansion[a][b][step] = add(p.expansion[a][b][step], map[rune]int{c: -1})
	}

	return p.expansion[a][b][step]
}

func add(c map[rune]int, d map[rune]int) map[rune]int {
	m := make(map[rune]int)

	for k, v := range c {
		m[k] += v
	}

	for k, v := range d {
		m[k] += v
	}

	return m
}

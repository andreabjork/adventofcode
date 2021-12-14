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
		makePolymer(ls, 30)
	}
}

func makePolymer(ls *bufio.Scanner, steps int) {
	STEPS := steps+1
	line, ok := util.Read(ls)
	// Create polymer
	templates := make([][]rune, STEPS)
	count := make(map[rune]int)
	templates[0] = []rune(line)
	for _, r := range templates[0] {
		count[r]++
	}
	p := Polymer{templates, count, make(map[rune]map[rune]rune)}

	// Record rules
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

	// Expand
	p.expand()

	min, max := int(^uint(0) >> 1), -1
	for _, v := range p.count {
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
	templates 	[][]rune
	count		map[rune]int
	rules	 	map[rune]map[rune]rune
}

func (p *Polymer) print(step int) {
	fmt.Println(string(p.templates[step]))
}

func (p *Polymer) expand() {
	for i, j := 0, 1; j < len(p.templates[0]); i,j = i+1, j+1 {
		p.insert(p.templates[0][i], p.templates[0][j], 0)
	}
}

func (p *Polymer) insert(a rune, b rune, steps int) {
	if steps == len(p.templates)-1 {
		return
	}

	steps++
	c := p.rules[a][b]

	p.count[c]++
	//p.track(a, b, c, steps)
	p.insert(a, c, steps)
	p.insert(c, b, steps)
}

func (p *Polymer) track(a rune, b rune, c rune, steps int) {
	p.count[c]++
	if len(p.templates[steps]) == 0 {
		p.templates[steps] = append(p.templates[steps], a, c, b)
	} else {
		p.templates[steps] = append(p.templates[steps][:len(p.templates[steps])-1], a, c, b)
	}
}

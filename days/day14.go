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
	count := make(map[rune]int)
	template := []rune(line)
	for _, r := range template {
		count[r]++
	}
	p := Polymer{steps, template, count, make(map[rune]map[rune]rune)}

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
	steps 		int
	template	[]rune
	count		map[rune]int
	rules	 	map[rune]map[rune]rune
}

func (p *Polymer) expand() {
	for i, j := 0, 1; j < len(p.template); i,j = i+1, j+1 {
		p.insert(p.template[i], p.template[j], 0)
	}
}

func (p *Polymer) insert(a rune, b rune, steps int) {
	if steps == p.steps {
		return
	}

	steps++
	c := p.rules[a][b]
	p.count[c]++
	p.insert(a, c, steps)
	p.insert(c, b, steps)
}

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
	p := Polymer{steps, template, count, make(map[rune]map[rune]rune), make(map[rune]map[rune][]Counter)}

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
	counter := p.expand()
	min, max := int(^uint(0) >> 1), -1
	for _, v := range counter.counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	fmt.Println("Most common - least common: ", max - min)
}

type Counter struct {
	counts	 	map[rune]int
}

func (c Counter) add(d Counter) Counter {
	m := make(map[rune]int)

	for k, v := range c.counts {
		m[k] += v
	}

	for k, v := range d.counts {
		m[k] += v
	}

	return Counter{m}
}

type Polymer struct {
	steps     int
	template  []rune
	counts    map[rune]int
	rules     map[rune]map[rune]rune
	expansion map[rune]map[rune][]Counter // expansion[x][y][3][z] = the number of z's added from XY after 3 steps
}

func (p *Polymer) expand() Counter {
	c := Counter{make(map[rune]int)}
	for i, j := 0, 1; j < len(p.template); i,j = i+1, j+1 {
		c = c.add(p.count(p.template[i], p.template[j], p.steps))

		if j < len(p.template) -1 {
			// remove duplicate count
			c = c.add(Counter{map[rune]int{p.template[j]: -1}})
		}
	}

	return c
}

func (p *Polymer) count(a rune, b rune, step int) Counter {
	// With xy -> z we get
	// expansion[x][y][3] = expansion[x][z][2]+expansion[z][y][2]
	c := p.rules[a][b]
	if step == 1 {
		cc := Counter{make(map[rune]int)}
		cc.counts[a]++
		cc.counts[b]++
		cc.counts[c]++
		return cc
	}

	// if XY expansion was never counted, create the counters
	if p.expansion[a] == nil {
		p.expansion[a] = make(map[rune][]Counter)
	}
	if p.expansion[a][b] == nil {
		p.expansion[a][b] = make([]Counter, p.steps+1)
	}

	// if XY expansion is not recorded at step 'step', count it now
	if p.expansion[a][b][step].counts == nil {
		p.expansion[a][b][step] = p.count(a, c, step-1).add(p.count(c, b, step-1).add(Counter{map[rune]int{c: -1}}))
	}

	return p.expansion[a][b][step]
}

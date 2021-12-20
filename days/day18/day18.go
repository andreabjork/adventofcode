package day18

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day18(inputFile string, part int) {
	if part == 0 {
		res, mag := FishmathA(inputFile)
		fmt.Println(res)
		fmt.Printf("Magnitude: %d\n", mag)
	} else {
		res, mag := FishmathB(inputFile)
		fmt.Println(res)
		fmt.Printf("Magnitude: %d\n", mag)
	}
}

func FishmathA(inputFile string) (string, int) {
	ls := util.LineScanner(inputFile)

	line, ok := util.Read(ls)
	var n, m *Noda
	n = parseTree(line)
	n.reduce()
	line, ok = util.Read(ls)
	for ok {
		m = parseTree(line)
		n = n.add(m)
		n.reduce()
		line, ok = util.Read(ls)
	}

	return n.print(), n.magnit()
}

func FishmathB(inputFile string) (string, int) {
	ls := util.LineScanner(inputFile)
	inputs := make([]string, 0)
	for line, ok := util.Read(ls); ok; {
		inputs = append(inputs, line)
		line, ok = util.Read(ls)
	}

	var maxN *Noda
	maxMag := 0
	for _, ele := range inputs {
		for _, toAdd := range inputs {
			n := parseTree(ele)
			m := parseTree(toAdd)
			n = n.add(m)
			n.reduce()
			if mag := n.magnit(); mag > maxMag {
				maxMag = mag
				maxN = n
			}
		}	
	}
	return maxN.print(), maxMag 
}

// =============
// DFS & METHODS
// -------------
// restart
// step
// explodes (if able)
// splits (if able)
type DFS struct {
	marked map[*Noda]bool
	root   *Noda
	n      *Noda
	depth  int
	str    string
}

func (dfs *DFS) restart() {
	dfs.marked = make(map[*Noda]bool)
	dfs.n = dfs.root
	dfs.depth = 0
	dfs.str = ""
}

func (dfs *DFS) step() bool {
	if dfs.n.left != nil && !dfs.marked[dfs.n.left] {
		dfs.marked[dfs.n] = true
		dfs.n = dfs.n.left
		dfs.depth++
	} else if dfs.n.right != nil && !dfs.marked[dfs.n.right] {
		dfs.marked[dfs.n] = true
		dfs.n = dfs.n.right
		dfs.depth++
	} else if !dfs.marked[dfs.n] {
		dfs.marked[dfs.n] = true
	} else {
		if dfs.n.parent == nil {
			return false
		}
		dfs.n = dfs.n.parent
		dfs.depth--
	}

	return true
}

func (dfs *DFS) explodes() bool {
	if dfs.depth >= 5 {
		dfs.n.parent.explode()
		return true
	}
	return false
}

func (dfs *DFS) splits() bool {
	if dfs.n.magnitude >= 10 && dfs.n.isLeaf() {
		var l, r int
		if dfs.n.magnitude%2 == 0 {
			l, r = dfs.n.magnitude/2, dfs.n.magnitude/2
		} else {
			l, r = dfs.n.magnitude/2, dfs.n.magnitude/2+1
		}
		dfs.n.magnitude = l + r
		dfs.n.left = &Noda{l, dfs.n, nil, nil}
		dfs.n.right = &Noda{r, dfs.n, nil, nil}
		return true
	}
	return false
}

// ===============
// NODA & METHODS
// ---------------
// - add
// - reduce
// - split
// - explode
// - firstRight
// - firstLeft
// - isLeft
// - isRight
// - isLeaf
// - magnit
// - print
// ===============
type Noda struct {
	magnitude int
	parent    *Noda
	left      *Noda
	right     *Noda
}

func (n *Noda) reduce() {
	root := n
	dfs := &DFS{make(map[*Noda]bool), root, n, 0, ""}
	ok := true
	reductions := 1
	for reductions > 0 {
		reductions = 0
		ok = true
		for ok {
			ok = dfs.step()
			if dfs.explodes() {
				dfs.restart()
				ok = true
			}

		}
		dfs.restart()
		ok = true
		for ok {
			ok = dfs.step()
			if dfs.splits() {
				ok = false
				dfs.restart()
				reductions++
			}
		}
	}
}

func (n *Noda) firstRight() (*Noda, bool) {
	m := n
	for m.isRight() {
		m = m.parent
		if m.parent == nil {
			return nil, false
		}
	}
	m = m.parent.right

	for !m.isLeaf() {
		m = m.left
	}
	return m, true
}

func (n *Noda) firstLeft() (*Noda, bool) {
	m := n
	for m.isLeft() {
		m = m.parent
		if m.parent == nil {
			return nil, false
		}
	}
	m = m.parent.left

	for !m.isLeaf() {
		m = m.right
	}
	return m, true
}

func (n *Noda) explode() {
	addLeft := n.left.magnitude
	addRight := n.right.magnitude
	if l, found := n.firstLeft(); found {
		l.magnitude += addLeft
	}

	if r, found := n.firstRight(); found {
		r.magnitude += addRight
	}
	n.magnitude = 0
	n.left = nil
	n.right = nil
}

func (n *Noda) add(m *Noda) *Noda {
	newRoot := &Noda{-1, nil, n, m}
	n.parent = newRoot
	m.parent = newRoot
	return newRoot
}

func (n *Noda) isLeft() bool {
	return n.parent.left == n
}

func (n *Noda) isRight() bool {
	return n.parent.right == n
}

func (n *Noda) isLeaf() bool {
	return n.right == nil && n.left == nil
}

func (n *Noda) magnit() int {
	if n.isLeaf() {
		return n.magnitude
	}
	return 3*n.left.magnit() + 2*n.right.magnit()
}

func (n *Noda) print() string {
	marked := make(map[*Noda]bool)
	str := ""
	for true {
		if n.left != nil && !marked[n.left] {
			str += "["
			marked[n] = true
			n = n.left
		} else if n.right != nil && !marked[n.right] {
			marked[n] = true
			n = n.right
		} else if !marked[n] {
			str += fmt.Sprintf("%d", n.magnitude)
			marked[n] = true
		} else {
			if n.parent == nil {
				break
			}
			if n == n.parent.right {
				str += "]"
			} else {
				str += ","
			}
			n = n.parent
		}
	}
	return str
}

// ============
// PARSE INPUTS
// ============
func parseTree(line string) *Noda {
	root := &Noda{-1, nil, nil, nil}
	n := root
	symbols := []rune(line)
	for _, sym := range symbols {
		if sym == '[' {
			n.left = &Noda{-1, n, nil, nil}
			n.right = &Noda{-1, n, nil, nil}
			n = n.left
		} else if sym == ']' {
			n = n.parent
		} else if sym == ',' {
			n = n.parent.right
		} else {
			v, _ := strconv.Atoi(string(sym))
			n.magnitude = v
		}
	}

	return root
}

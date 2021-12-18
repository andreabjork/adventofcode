package days

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
	"strconv"
)

func Day18(inputFile string, part int) {
	ls := util.LineScanner(inputFile)

	if part == 0 {
		doMath(ls)
	} else {
		fmt.Printf("bla")
	}
}

func parseInput(line string) *MTree {
	fmt.Println("Creating tree")
	node := &MNode{0, nil, nil, nil}
	tree := &MTree{nil, node, 0, 0,  0}

	addLeft := true
	symbols := []rune(line)
	for _, sym := range symbols {
		if sym == '[' {
			if tree.root == nil {
				fmt.Println("adding root")
				tree.root = node
			} else {
				fmt.Println("left")
				node.left = &MNode{-1, node, nil, nil}
			}
			tree.totalNodes++
		} else if sym == ']' {
			fmt.Println("Pivot up")
			if node.parent == nil {
				break
			}
			node = node.parent
			node.magnitude = 3*node.left.magnitude + 2*node.right.magnitude
		} else if sym == ',' {
			fmt.Println("Pivot right")
			addLeft = false
		} else {
			val, _ := strconv.Atoi(string(sym))
			if addLeft {
				fmt.Println("left literal ", val)
				node.left = &MNode{val, node, nil, nil}
				tree.totalNodes++
			} else {
				fmt.Println("right literal ", val)
				node.right = &MNode{val, node, nil, nil}
				addLeft = true
				tree.totalNodes++
			}
		}
	}

	return tree
}

func doMath(ls *bufio.Scanner) {
	line, ok := util.Read(ls)

	tree := parseInput(line)
	tree.print()
	tree.reduce()
	tree.print()
	for ok {
		ttree := parseInput(line)
		tree.add(ttree)
		tree.print()
		tree.reduce()
		tree.print()

		line, ok = util.Read(ls)
	}
	fmt.Println(tree.root.magnitude)
}

type MTree struct {
	root        *MNode
	curr 		*MNode
	height 		int
	node 		int
	totalNodes 	int
}

type MNode struct {
	magnitude int
	parent    *MNode
	left      *MNode
	right 	  *MNode
}

func (t *MTree) print() {
	for t.node < t.totalNodes {
		if t.curr.left != nil {
			t.curr = t.curr.left
			t.height++
			fmt.Printf("[")
		} else if t.curr.right != nil {
			t.curr = t.curr.right
			t.height++
			fmt.Printf("[")
		} else {
			fmt.Printf("%d", t.curr.magnitude)
			t.curr = t.curr.parent
			t.height--
		}

		t.node++
	}

}

func (t *MTree) reduce() {
	reduced := t.reduceOne()
	for reduced {
		reduced = t.reduceOne()
	}
}

func (t *MTree) reduceOne() bool {
	reduced := false
	for t.node < t.totalNodes {
		if t.curr.left != nil {
			t.curr = t.curr.left
			t.height++
		} else if t.curr.right != nil {
			t.curr = t.curr.right
			t.height++
		} else {
			t.curr = t.curr.parent
			t.height--
		}

		t.node++
		if t.height >= 5 {
			t.explode()
			reduced = true
			break
		} else if t.curr.magnitude >= 10 {
			t.split()
			reduced = true
			break
		}
	}

	return reduced
}

func (t *MTree) explode() {
	fmt.Println("Exploding!")
	if t.curr.parent.left == t.curr {
		t.curr.parent.right.magnitude += t.curr.magnitude
	} else {
		t.curr.parent.left.magnitude += t.curr.magnitude
	}
	t.curr = &MNode{0, t.curr.parent, nil, nil}
}

func (t *MTree) split() {
	fmt.Println("Splitting!")
	var l, r int
	if t.curr.magnitude % 2 == 0 {
		l, r = t.curr.magnitude/2, t.curr.magnitude/2
	} else {
		l, r = t.curr.magnitude/2, t.curr.magnitude/2+1
	}

	t.curr.magnitude = l+r
	t.curr.left = &MNode{l, t.curr, nil, nil}
	t.curr.right = &MNode{r, t.curr, nil, nil}
}

func (t *MTree) add(u *MTree) {
	t.root = &MNode{3*t.root.magnitude + 2*u.root.magnitude, nil, t.root, u.root}
	t.height++
	t.totalNodes += u.totalNodes
}
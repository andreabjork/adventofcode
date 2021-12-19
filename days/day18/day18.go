package days

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day18(inputFile string, part int) {
	if part == 0 {
		Fishmath(inputFile)
	} else {
		fmt.Printf("bla")
	}
}

func Fishmath(inputFile string) (string, int) {
	ls := util.LineScanner(inputFile)

	line, ok := util.Read(ls)
	var str string
	var root, n, m *Noda
	n = parseTree(line)
	root = n
	str = print(n)
	fmt.Println(str)
	line, ok = util.Read(ls)
	for ok {
		m = parseTree(line)
		strm := print(m)
		fmt.Printf("adding %s to %s\n", strm, str)
		n = n.add(m)

		str = print(n)
		fmt.Println("After add:")
		fmt.Println(str)

		n.reduce()

		str = print(n)
		fmt.Println("After reduce:")
		fmt.Println(str)

		line, ok = util.Read(ls)
	}

	return str, root.magnitude
}

type Noda struct {
	magnitude 	int
	parent    	*Noda
	left      	*Noda
	right     	*Noda
}

func (n *Noda) firstRight() (*Noda, bool) {
	m := n
	if m.isLeft() {
		fmt.Println("it's a left node")
		m = m.parent.right
	} else if m.isRight() {
		fmt.Println("its a right node")
		m = m.parent.parent
		for m.parent != nil && m.isRight() {
			m = m.parent
		}

		if m.parent == nil {
			// reached root
			fmt.Println("No right node!")
			return nil, false
		}
		m = m.right
	}

	fmt.Println("Starting at ", m)
	for !m.isLeaf() {
		m = m.left
	}
	return m, true
}

func (n *Noda) firstLeft() (*Noda, bool) {
	m := n
	if m.isRight() {
		fmt.Println("it's a right node")
		m = m.parent.left
	} else if m.isLeft() {
		fmt.Println("its a left node")
		m = m.parent.parent
		for m.parent != nil && m.isLeft() {
			m = m.parent
		}
		if m.parent == nil {
			// reached root
			fmt.Println("No left node!")
			return nil, false
		}
		m = m.left
	}


	fmt.Println("Starting at ", m)
	for !m.isLeaf() {
		m = m.right
	}
	return m, true
}

func (n *Noda) explode() {
	fmt.Printf("Exploding %d %d\n", n.left.magnitude, n.right.magnitude)
	addLeft := n.left.magnitude
	addRight := n.right.magnitude


	fmt.Println("finding first left")
	if l, found := n.firstLeft(); found {
		fmt.Printf("Adding %d to %d\n", addLeft, l.magnitude)
		l.magnitude += addLeft
	}

	fmt.Println("finding first right")
	if r, found := n.firstRight(); found {
		fmt.Printf("Adding %d to %d\n", addRight, r.magnitude)
		r.magnitude += addRight
	}
	n.magnitude = 0
	n.left = nil
	n.right = nil

}

func (n *Noda) split() {
	var l, r int
	if n.magnitude % 2 == 0 {
		l, r = n.magnitude/2, n.magnitude/2
	} else {
		l, r = n.magnitude/2, n.magnitude/2+1
	}
	n.magnitude = l+r
	n.left = &Noda{l, n, nil, nil}
	n.right = &Noda{r, n, nil, nil}
}

func (n *Noda) add(m *Noda) *Noda {
	newRoot := &Noda{3*n.magnitude + 2*n.magnitude, nil, n, m}
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

// ----------------
func (n *Noda) reduce() {
	marked := make(map[*Noda]bool)
	root := n
	depth := 0
	str := ""
	for true {
		if n.left != nil && !marked[n.left] {
			marked[n] = true
			n = n.left
			depth++
		} else if n.right != nil && !marked[n.right] {
			marked[n] = true
			n = n.right
			depth++
		} else if !marked[n] {
			str += fmt.Sprintf("%d", n.magnitude )
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
			depth--
		}

		if depth >= 5 {
			n.parent.explode()
			fmt.Println("After explode: ")
			str := print(root)
			fmt.Println(str)
			n = root
			marked = make(map[*Noda]bool)
			depth = 0
		} else if n.magnitude >= 10 && n.isLeaf() {
			n.split()
			fmt.Println("After split: ")
			str := print(root)
			fmt.Println(str)
			n = root
			marked = make(map[*Noda]bool)
			depth = 0
		}
	}
}

func print(n *Noda) string {
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
			str += fmt.Sprintf("%d", n.magnitude )
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

func parseTree(line string) *Noda {
	root := &Noda{0, nil, nil, nil}
	n := root
	symbols := []rune(line)
	for _, sym := range symbols {
		if sym == '[' {
			n.left = &Noda{0, n, nil, nil}
			n.right = &Noda{0, n, nil, nil}
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

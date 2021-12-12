package days

import (
	"adventofcode/m/v2/util"
	"fmt"
	"sort"
)

func Day10(inputFile string, part int) {
	if part == 0 {
		day10a(inputFile)
	} else {
		day10b(inputFile)
	}
}

func day10a(inputFile string) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	var score int
	sum := 0
	for ok {
		_, score = checkSyntax([]rune(line))
		sum += score
		line, ok = util.Read(ls)
	}

	fmt.Printf("Syntax err errScore: %d\n", sum)
}

func day10b(inputFile string) {
	ls := util.LineScanner(inputFile)
	line, ok := util.Read(ls)

	var stack *Stack
	var score int
	scores := make([]int, 0)
	for ok {
		stack, score = checkSyntax([]rune(line))
		// Only do completion scores for lines that are ok
		if score == 0 {
			scores = append(scores, completionScore(stack))
		}
		line, ok = util.Read(ls)
	}

	sort.Ints(scores)
	fmt.Printf("Completion score: %d\n", scores[len(scores)/2])
}

var delimiters = map[rune]*Delim {
	'(': {true, ')', 1 },
	'[': {true, ']', 2 },
	'{': {true, '}', 3 },
	'<': {true, '>', 4 },
	')': {false, '(', 3 },
	']': {false, '[', 57},
	'}': {false, '{', 1197},
	'>': {false, '<', 25137},
}

// Returns the remaining stack, and the error score
func checkSyntax(runes []rune) (*Stack, int) {
	stack := Stack{nil }
	for _, r := range runes {
		d := delimiters[r]
		if d.opening {
			stack.push(d)
		} else {
			p, ok := stack.pop()
			if !ok {
				fmt.Println("Received closing delimiter on empty stack.")
				break
			} else if !p.closes(r) {
				// fmt.Printf("Syntax error: Expected %s, found %s \n", string((*p).matcher), string(r))
				return &stack, d.score
			}
		}
	}

	return &stack, 0
}

// Calculates the completion score from the remaining stack
func completionScore(s *Stack) int {
	total := 0
	// Pop the rest of the stack and count the matrix
	for p, ok := (*s).pop(); ok; {
		total = 5*total + (*p).score
		p, ok = (*s).pop()
	}

	return total
}

// ==========
// DELIMITERS
// ==========
type Delim struct {
	opening    bool
	matcher    rune
	score      int
}

func (d *Delim) closes(r rune) bool {
	return (*d).matcher == r
}

// ====================
// STACK IMPLEMENTATION
// ====================
type StackEle struct {
	val		*Delim
	next	*StackEle
}

type Stack struct {
	top 	*StackEle
}

func (s *Stack) push(d *Delim) {
	(*s).top = &StackEle{d, (*s).top}
}

func (s *Stack) pop() (*Delim, bool) {
	if (*s).top == nil {
		return nil, false
	}

	d := (*s).top.val
	(*s).top = (*s).top.next

	return d, true
}

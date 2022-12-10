package day23

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
)

func Day23(inputFile string, part int) {
	if part == 0 {
		min := Play(inputFile, 2)
		fmt.Printf("Minimum fuel: %d\n", min)
	} else {
		min := Play(inputFile, 4)
		fmt.Printf("Minimum fuel: %d\n", min)
	}
}

func Play(inputFile string, levels int ) int {
	const MAX_INT = int(^uint(0) >> 1)
	start := parseState(inputFile, levels)
	game := &Game{MAX_INT, nil, start}
	game.decisionTree()
	// Show minimum game
	s := game.minState
	for s != nil {
		s.show()
		s = s.fromState
	}

	return game.minCost
}

// ====
// GAME
// ====
type Game struct {
	minCost 		int
	minState 		*State
	currentState 	*State
}

// ====================
// STATE: DECISION TREE
// ====================
// Decision Tree with nodes containing 'State'
type State struct {
	score 						int //
	board 						[][]rune
	amphipods					[]*Amphipod
	fuel						int
	fromState 					*State
	nextStates					[]*State
}

func (g *Game) decisionTree() int {
	g.DFS(g.currentState)
	return g.minCost
}

// Depth First Search of the optimum, constructing the
// tree as we go
func (g *Game) DFS(s *State) {
	//s.show()
	if s.score == 0 {
		//fmt.Println("Winning state!")
		if s.fuel <= g.minCost {
			g.minCost = s.fuel
			g.minState = s
		}
	} else if s.score >= 0 && s.fuel >= g.minCost {
		// exceeded known minimum, stop
		g.currentState = s.fromState
	} else if s.score >= 0 {
		s.addNextStates()
		for _, st := range s.nextStates {
			g.DFS(st)
		}
	}
}

// Next possible states:
// if any amphipod can go home, this will be the only possible next state.
//
// Any unblocked amphipod moves to any* possible position,
// so long as that position makes sense.
func (s *State) addNextStates() {
	s.nextStates = []*State{}
	for amp, a := range s.amphipods {
		// Record possible next moves for all unblocked amphipods
		if yes, floor := s.canGoHome(a); yes {
			s.nextStates = []*State{s.toHome(amp, floor)}
			break
		} else if yes, space := s.canGoHallway(a); yes {
			for i := 0; i < len(space); i++ {
				s.nextStates = append(s.nextStates, s.toHallway(amp, space[i]))
			}
		}
	}
}

// All possible moves:
//  - Moving directly home
//  - Moving to hallway, never in front of a room
func (s *State) toHome(amp int, floor int) *State {
	ss := s.createFrom()
	ss.score--
	a := ss.amphipods[amp]
	ss.fuel = ss.fuel + ((a.vpos+floor)*a.energy())+(util.Abs(a.home()-a.hpos)*a.energy())
	ss.board[a.vpos][a.hpos] = '.'
	a.hpos = a.home()
	a.vpos = floor
	ss.board[a.vpos][a.hpos] = a.subtype
	return ss
}

func (s *State) toHallway(amp int, i int) *State {
	ss := s.createFrom()
	a := ss.amphipods[amp]
	ss.fuel = ss.fuel + (a.vpos*a.energy())+(util.Abs(i-a.hpos)*a.energy())
	//fmt.Printf("Sending %s from %d,%d to Hallway at %d for %d fuel\n", string(a.subtype), a.vpos, a.hpos, i, fuelSpent)
	ss.board[a.vpos][a.hpos] = '.'
	a.hpos = i
	a.vpos = 0
	ss.board[a.vpos][a.hpos] = a.subtype
	return ss
}

func (s *State) createFrom() *State {
	ss := &State{
		s.score,
		[][]rune{},
		[]*Amphipod{},
		s.fuel,
		s,
		[]*State{},
	}

	for i := 0; i < len(s.board); i++ {
		ss.board = append(ss.board, []rune{})
		for j := 0; j < len(s.board[i]); j++ {
			ss.board[i] = append(ss.board[i], s.board[i][j])
		}
	}

	for i := 0; i < len(s.amphipods); i++ {
		amp := &Amphipod{s.amphipods[i].subtype, s.amphipods[i].vpos, s.amphipods[i].hpos}
		ss.amphipods = append(ss.amphipods, amp)
	}
	return ss
}


func (s *State) show() {
	fmt.Printf("Score: %d\n", s.score)
	fmt.Printf("Fuel: %d\n", s.fuel)
	fmt.Printf("#############\n#")
	for i := 0; i < len(s.board[0]); i++ {
		fmt.Printf("%s", string(s.board[0][i]))
	}
	fmt.Printf("#\n##")
	spaces := []int{2,4,6,8}
	for i := 0; i < len(spaces); i++ {
		x := spaces[i]
		fmt.Printf("#%s", string(s.board[1][x]))
	}
	fmt.Printf("###\n")

	for l := 2; l < len(s.board); l++ {
		fmt.Printf("  ")
		for i := 0; i < len(spaces); i++ {
			x := spaces[i]
			fmt.Printf("#%s", string(s.board[l][x]))
		}
		fmt.Printf("#\n")
	}
	fmt.Printf("  #########\n\n")
}

// ========
// Amphipod
// ========
type Amphipod struct {
	subtype 	rune
	vpos 		int // 0 for being in room, 1 for hallway
	hpos 		int // number of room, or number of spot in hallway counting from left to right
}

func (a *Amphipod) energy() int {
	return map[rune]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}[a.subtype]
}

func (s *State) blockedInRoom(a *Amphipod) bool {
	// Everyone below (higher l) a must be the same type
	for l := a.vpos-1; l > 0; l-- {
		if s.board[l][a.hpos] != '.' {
			return true
		}
	}
	return false
}

func (s *State) atHome(a *Amphipod) bool {
	if a.hpos == a.home() {
		// Everyone below (higher l) a must be the same type
		for l := a.vpos+1; l < len(s.board); l++ {
			if ! (s.board[l][a.home()] == a.subtype) {
				return false
			}
		}
		return true
	}
	return false
}

func (a *Amphipod) home() int{
	return map[rune]int{'A': 2, 'B': 4, 'C': 6, 'D': 8}[a.subtype]
}

func (s *State) canGoHome(a *Amphipod) (bool, int) {
	if s.atHome(a) || s.blockedInRoom(a) {
		return false, 0
	}

	for l := len(s.board)-1; l > 0; l-- {
		if s.board[l][a.home()] == '.' {
			return s.hallwayFree(a.hpos, a.home()), l
		} else if s.board[l][a.home()] != a.subtype {
			return false, -1
		}
	}
	return false, -1
}

func (s *State) canGoHallway(a *Amphipod) (bool, []int) {
	if s.atHome(a) || a.vpos == 0 || s.blockedInRoom(a) {
		return false, []int{}
	}

	okPos := []int{}
	// Check all potential hallway positions then
	for _, hpos := range []int{0,1,3,5,7,9,10} {
		ok := s.hallwayFree(a.hpos, hpos)
		if ok {
			okPos = append(okPos, hpos)
		}
	}

	return len(okPos) > 0, okPos
}

func (s *State) hallwayFree(from int, to int) bool {
	free := true
	// from = 6, to 5 : i = 7, i < 5 NOPE
	// i = 6-1 = 5, i >= 5
	// Hallway path + hallway spot need to be clear
	// To the right:
	for i := from+1; i <= to; i++ {
		free = free && s.board[0][i] == '.'
	}
	// To the left:
	for i := from-1; i >= to; i-- {
		free = free && s.board[0][i] == '.'
	}

	return free
}

func createAmphipod(r rune, f int, i int) *Amphipod {
	return &Amphipod {
		r,
		f,
		i,
	}
}

// The boring stuff: Parsing input. This can be hard coded anyway but lets support different starting states.
func parseState(inputFile string, levels int) *State{
	ls := util.LineScanner(inputFile)
	_, _ = util.Read(ls)
	_, _ = util.Read(ls)
	board := [][]rune{
		{'.','.','.','.','.','.','.','.','.','.','.'},
	}

	for l := 0; l < levels; l++ {
		board = append(board, make([]rune, 4))
	}

	amphipods := []*Amphipod{}
	// top rooms
	line, _ := util.Read(ls)
	re := regexp.MustCompile("#+([A-D])#([A-D])#([A-D])#([A-D])#+")
	score := 0
	for l := 1; l <= levels; l++ {
		trunes := re.FindStringSubmatch(line)
		board[l] = make([]rune,11)
		spaces := []int{2,4,6,8}
		for i := 0; i < len(spaces); i++ {
			x := spaces[i]
			board[l][x] = []rune(trunes[i+1])[0]
			a := createAmphipod(board[l][x], l, x)
			amphipods = append(amphipods, a)
		}
		line, _ = util.Read(ls)
	}

	s := &State{
		score,
		board,
		amphipods,
		0,
		nil,
		[]*State{},
	}
	initialScore := 0
	for _, a := range s.amphipods {
		if !s.atHome(a) {
			initialScore++
		}
	}
	s.score = initialScore
	return s
}


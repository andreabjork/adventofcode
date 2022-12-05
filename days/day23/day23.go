package day23

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
)

func Day23(inputFile string, part int) {
	fmt.Println("not implemented")
}

func parseState(inputFile string) *State{
	// ###B#C#B#D###
  //   #A#D#C#A#
	ls := util.LineScanner(inputFile)
	//#############
	//#...........#
	//###B#C#B#D###
	// 	#A#D#C#A#
	// 	#########
	_, _ = util.Read(ls)
	_, ok := util.Read(ls)
	
	top, _ := util.Read(ls)
	bottom, _ := util.Read(ls)
	re := regexp.MustCompile(`#+([A-D])#([A-D])#([A-D])$([A-D])#+`)
	
	for i := 0; i < 2; i++ {
		line, _ := util.Read(ls)
		aps := re.FindStringSubmatch(line)

		for j := 0; j < len(aps); j++ {

		}

	}
	re.FindStringSubmatch(bottom)
	
	

	return &State{
		score,
		amphipods,
		0,
		nil,
		[]*State{},
	}
}

// ====
// GAME
// ====
type Game struct {
	minCost 			int
	currentState 	*State
}

func optimize() {
	start := parseState()
	game := &Game(MAX_INT, start)	
} 

// ====================
// STATE: DECISION TREE
// ====================
// Decision Tree with nodes containing 'State'
type State struct {
	score 							int // 
	amphipods						[]*Amphipod
	fuel								int
	fromState 					*State
	nextStates					[]*State
}


// Depth First Search of the optimum, constructing the
// tree as we go
func (s *State) DFS(fromNode *State) {
	if s.score == 0 {
		if s.fuel <= min {
			min = s.fuel
		}
	}

	// if score == 0, stop and track minimum
	// if score >= 0 but we've exceeded known minimum, stop
	// if score >= 0 and we've not exceeded known minimum, add next states

	// Next possible states: 
	// if any amphipod can go home, this will be the only possible next state.
	// 
	// Any unblocked amphipod moves to any* possible position,
	// so long as that position makes sense.
}

// Add the next states in the tree, taking care
// to limit states reasonably
func (s *State) addNextStates() {

}

func (s *State) show() {
	// print current gamestate
}

// ========
// Amphipod
// ========
type Amphipod struct {
	vpos int // 0 for being in room, 1 for hallway
	hpos int // number of room, or number of spot in hallway counting from left to right
}

func (a *Amphipod) walk() {
	
}

func (a *Amphipod) isBlocked() {

}

func (a *Amphipod) canGoHome() {

}
package day23

import (
	"fmt"
)

func Day23(inputFile string, part int) {
	fmt.Println("not implemented")
}

// ====================
// STATE: DECISION TREE
// ====================
// Decision Tree with nodes containing 'State'
type State struct {
	wanderingAmphipods	int
	amphipods						[]*Amphipod
	fuel								int
	fromState 					*State
	nextStates					[]*State
}

// Depth First Search of the optimum, constructing the
// tree as we go
func (s *State) DFS(fromNode *State) {
 	
}

// Add the next states in the tree, taking care
// to limit states reasonably
func (s *State) addNextStates() {

}

// ========
// Amphipod
// ========
type Amphipod struct {
	vpos int // 0 for being in room, 1 for hallway
	hpos int // number of room, or number of spot in hallway counting from left to right
}

func (a *Amphipod) canGoHome() {

}
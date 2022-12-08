package day23

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
)

func Day23(inputFile string, part int) {
	fmt.Println("not implemented")
}

func (s *State) print() {
	fmt.Printf("#############\n")
	fmt.Printf("#")
	for i := 0; i < len(s.hallway); i++ {
		fmt.Printf("%s", string(s.hallway[i]))
	}
	fmt.Printf("\n ")
	for i := 0; i < len(s.topRooms); i++ {
		fmt.Printf("#%s", string(s.topRooms[i]))
	}

	fmt.Printf("\n ")
	for i := 0; i < len(s.bottomRooms); i++ {
		fmt.Printf("#%s", string(s.topRooms[i]))
	}
	fmt.Printf("#########")
}

func parseState(inputFile string) *State{
	ls := util.LineScanner(inputFile)
	//#############
	//#...........#
	//###B#C#B#D###
	// 	#A#D#C#A#
	// 	#########
	_, _ = util.Read(ls)
	_, ok := util.Read(ls)

	hallway := make([]rune, 11)
	top, _ := util.Read(ls)
	bottom, _ := util.Read(ls)
	re := regexp.MustCompile(`#+([A-D])#([A-D])#([A-D])$([A-D])#+`)
	topRooms := re.FindStringSubmatch(top)	
	bottomRooms := re.FindStringSubmatch(bottom)	
	amphipods := make([]*Amphipod, 8)
	energies := map[rune]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}
	
	for top := amphipod
	amphipods[i*j] = &Amphipod{rType, energies[rType], i+1,j}	
	re.FindStringSubmatch(bottom)
	
	// Return the initial game state
	return &State{
		score,
		amphipods,
		hallway,
		topRooms,
		bottomRooms,
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
	hallway 						[]rune
	topRooms						[]rune
	bottomRooms					[]rune
	amphipods						[]*Amphipod
	fuel								int
	fromState 					*State
	nextStates					[]*State
}


// Depth First Search of the optimum, constructing the
// tree as we go
//func (s *State) DFS(fromNode *State) {
//	if s.score == 0 {
//		if s.fuel <= min {
//			min = s.fuel
//		}
//	}

	// if score == 0, stop and track minimum
	// if score >= 0 but we've exceeded known minimum, stop
	// if score >= 0 and we've not exceeded known minimum, add next states

	// Next possible states: 
	// if any amphipod can go home, this will be the only possible next state.
	// 
	// Any unblocked amphipod moves to any* possible position,
	// so long as that position makes sense.
//}

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
	subtype rune
	energy  int
	vpos 		int // 0 for being in room, 1 for hallway
	hpos 		int // number of room, or number of spot in hallway counting from left to right
}

func (a *Amphipod) getEnergy() rune {
	return []rune{'A': 1, 'B': 10, 'C': 100, 'D': 1000}[a.subtype]
}

//func (a *Amphipod) walk() {
//	
//}
//
//func (a *Amphipod) isBlocked() {
//
//}
//
//func (a *Amphipod) canGoHome() {
//
//}
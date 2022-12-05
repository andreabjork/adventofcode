package day23

import (
	"fmt"
)

func Day23(inputFile string, part int) {
	fmt.Println("not implemented")
}

// Decision Tree with nodes containing 'State'
type State struct {
	wanderingAmphipods	int
	amphipods						[]*Amphipod
	fuel								int
}

// Depth First Search of the optimum, constructing the
// tree as we go
func DFS(fromNode *State) {
	
}

package day23

import (
	"fmt"
)

func Day23(inputFile string, part int) {
	fmt.Println("not implemented")
}

// Solution:
// 
// We will define a game.getHome(a *Amphipod) action which
// aims to get the amphipod a to its home cell, moving
// all other amphipods out of the way first.
//
// We iterate over amphipods and get each one home while
// tallying up the costs. If a cost exceeds the existing
// minima then we stop, and try the next.
//
// Costs will
func PlayAmphipods(inputFile string) int{

	return 0
}

type Tile int64
const (
	Wall Tile = iota
	Hallway
	RoomA
	RoomB
	RoomC
	RoomD
)

type Game struct {
	cost 	int
	board 	map[int]map[int]Tile
}

type Amphipod struct {
	stepCost	int // 
	home 		int // room 1,2,3,4
	location 	int // room 1,2,3,4 or 0 for hallway
}

func (a *Amphipod) goHallway() {

}

func (a *Amphipod) goHome() {

}

package day21

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
)

func Day21(inputFile string, part int) {
	if part == 0 {
		res := PlayDeterministic(inputFile)
		fmt.Println("Result: ", res)
	} else {
		winner := PlayDirac(inputFile)
		fmt.Println("Most wins: ", winner)
	}
}
func PlayDeterministic(inputFile string) int {
	ls := util.LineScanner(inputFile)
	line, _ := util.Read(ls)
	p1 := strings.Split(line, "Player 1 starting position: ")[1]
	line, _ = util.Read(ls)
	p2 := strings.Split(line, "Player 2 starting position: ")[1]

	player1 := &Player{0, util.ToInt(p1)}
	player2 := &Player{0, util.ToInt(p2)}

	players := []*Player{player1, player2}
	die := &DD{0, 0}
	r := 0
	for !player1.won() && !player2.won() {
		players[r].advance(die.roll() + die.roll() + die.roll())
		r = (r + 1) % 2
	}

	var result int
	if player1.won() {
		result = player2.score * die.rolls
	} else {
		result = player1.score * die.rolls
	}

	return result
}

type Player struct {
	score int
	track int
}

func (p *Player) advance(n int) {
	p.track = (p.track+n-1)%10 + 1
	p.score += p.track
}

func (p *Player) won() bool {
	return p.score >= 1000
}

type DD struct {
	value int
	rolls int
}

func (d *DD) roll() int {
	d.value = (d.value % 100) + 1
	d.rolls++
	return d.value
}

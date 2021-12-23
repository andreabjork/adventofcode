package day21

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
)

// From permutations of a,b,c with a,b,c in {1,2,3}
// Each 3-roll will give you all combinations, which is a score of either 3, 4, 5, 6, 7 or 9.
// You get this set of results:
// [3 4 4 4 5 5 5 5 5 5 6 6 6 6 6 6 6 7 7 7 7 7 7 8 8 8 9]
var outcomes = map[int]int64 {
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

// States is a DP matrix containing
// map[p1][p1score][p2][p2score] = p1wins,p2wins
var states = make(map[int]map[int]map[int]map[int][]int64)

func PlayDirac(inputFile string) int64 {
	ls := util.LineScanner(inputFile)
	line, _ := util.Read(ls)
	player1 := strings.Split(line, "Player 1 starting position: ")[1]
	line, _ = util.Read(ls)
	player2 := strings.Split(line, "Player 2 starting position: ")[1]

	p1 := util.ToInt(player1)
	p2 := util.ToInt(player2)
	p1score, p2score := 0,0
	p1wins, p2wins := Wins(p1, p1score, p2, p2score)

	fmt.Println("Player 1: ", p1wins)
	fmt.Println("Player 2: ", p2wins)

	if p1wins > p2wins {
		return p1wins
	} else {
		return p2wins
	}
}

// DP Recursion calculating the number of wins for both players, which results from the
// next roll in 'outcomes' with player 1 and 2 being at p1, p2 on the track respectively,
// and having accumulated a score of p1score, p2score respectively.
func Wins(p1 int, p1score int, p2 int, p2score int) (int64, int64) {
	if p1score >= 21 {
		return 1, 0
	} else if p2score >= 21 {
		return 0, 1
	}

	if val, found := states[p1][p2][p1score][p2score]; found {
		return val[0], val[1]
	}

	var w1, w2 int64
	for roll, universes := range outcomes {
		next_p1 := (p1+roll-1)%10+1
		next_p1score := p1score+next_p1
		ww2, ww1 := Wins(p2, p2score, next_p1, next_p1score)
		w1, w2 = w1+ww1*universes, w2+ww2*universes
	}

	if states[p1] == nil {
		states[p1] = make(map[int]map[int]map[int][]int64)
	} 
	if states[p1][p2] == nil {
		states[p1][p2] = make(map[int]map[int][]int64)
	}
	if states[p1][p2][p1score] == nil {
		states[p1][p2][p1score] = make(map[int][]int64)
	} 

	states[p1][p2][p1score][p2score] = []int64{w1, w2}
	return w1, w2
}

package day6

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
)


func Day6(inputFile string, part int) {
	if part == 0 {
		day6(inputFile, 80)
	} else {
		day6(inputFile, 256)
	}
}

func day6(inputFile string, N int) {
	ws := util.LineScanner(inputFile)
	read, _ := util.Read(ws)

	initState := strings.Split(read, ",")

	// Initialize due date slice
	var (
		M = 10 // fish born on day 0 will give birth on day 9, so we need to track 10-day windows
		dueDates = make([]int, M) // [ x0 x1 x2 ... ], where xi = #fish due on day n = i (mod M)
		dueDate	int
		fish = 0
		today int // day number mod M
		newFish int // total of new fish born today
		newFishDD int // expected due date of new fish, mod M
		oldFishDD int // expected due date of old fish, mod M
	)
	for _, ini := range initState {
		// fish labelled 0 produces on day 1
		dueDate = util.ToInt(ini) + 1
		dueDates[dueDate]++
		fish++
	}

	// simulate births for N days
	for day := 1; day <= N; day++ {
		today = day % M
		// birth new fish
		newFish = dueDates[today]
		fish += newFish

		// update due dates
		newFishDD = (today +9)%M
		oldFishDD = (today +7)%M
		dueDates[newFishDD] += newFish
		dueDates[oldFishDD] += newFish
		dueDates[today] -= newFish
	}

	fmt.Printf("# Fish @ Day %d: %d\n", N, fish)
}

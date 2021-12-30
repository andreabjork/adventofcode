package day7

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
)


func Day7(inputFile string, part int) {
	if part == 0 {
		day7(inputFile, da)
	} else {
		day7(inputFile, db)
	}
}

func day7(inputFile string, d func(m int, x int) int) {
	ws := util.LineScanner(inputFile)
	read, _ := util.Read(ws)

	positions := strings.Split(read, ",")
	pos := util.ToInt(positions[0])
	X := make([]int, 1)
	X[0] = pos
	minM := pos
	maxM := pos
	const MaxUint = ^uint(0)
	minSum := int(MaxUint >> 1)

	for i := 1; i < len(positions); i++ {
		x := util.ToInt(positions[i])
		X = append(X, x)
		if x < minM {
			minM = x
		}
		if x > minM {
			maxM = x
		}
	}

	// Could search for the m in a much more optimal way but ¯\_(ツ)_/¯
	for m := minM; m <= maxM; m++ {
		s := 0
		for _, x := range X {
			s += d(m, x)
		}
		if abs(s) <= minSum {
			minSum = abs(s)
		} else {
			break
		}
	}

	fmt.Printf("Minimum fuel: %d\n", minSum)
}

// Distance metric for part a: you can also just cheat this iteration by calculating median directly
func da(m int, x int) int {
	return abs(m-x)
}

// Distance metric for part b
func db(m int, x int) int {
	return abs(m - x)*(abs(m-x)+1)/2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

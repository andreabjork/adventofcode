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

func inCuboid(x,y,z,)

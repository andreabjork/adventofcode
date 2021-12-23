package main

import (
	"adventofcode/m/v2/days"
	"adventofcode/m/v2/days/day18"
	"adventofcode/m/v2/days/day19"
	"adventofcode/m/v2/days/day21"
	"adventofcode/m/v2/util"
	"fmt"
	"os"
	"time"
)

var dayfuncs = map[int]interface{}{
	1:  days.Day1,
	2:  days.Day2,
	3:  days.Day3,
	4:  days.Day4,
	5:  days.Day5,
	6:  days.Day6,
	7:  days.Day7,
	8:  days.Day8,
	9:  days.Day9,
	10: days.Day10,
	11: days.Day11,
	12: days.Day12,
	13: days.Day13,
	14: days.Day14,
	15: days.Day15,
	16: days.Day16,
	17: days.Day17,
	18: day18.Day18,
	19: day19.Day19,
	21: day21.Day21,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Run as:\n\tgo run main.go all\nor as:\n\tgo run main.go [day] [0|1] [test|real]")
		return
	}

	var (
		day        int
		part       int
		test_input string
		real_input string
	)
	if os.Args[1] == "all" {
		for d := 1; d <= 18; d++ {
			test_input = fmt.Sprintf("inputs/day%d_test.txt", d)
			real_input = fmt.Sprintf("inputs/day%d.txt", d)
			runDay(d, 0, test_input)
			runDay(d, 0, real_input)
			runDay(d, 1, test_input)
			runDay(d, 1, real_input)
		}
	} else {
		// Otherwise run the specified day
		day = util.ToInt(os.Args[1])
		part = util.ToInt(os.Args[2])
		if len(os.Args) > 3 {
			test_input = os.Args[3] //fmt.Sprintf("inputs/day%d_test.txt", day)
			runDay(day, part, test_input)
		} else {
			real_input = fmt.Sprintf("inputs/day%d.txt", day)
			runDay(day, part, real_input)
		}
	}
}

func runDay(day int, part int, input string) {
	fmt.Printf("Day %d:%d < %s\n", day, part, input)
	start := time.Now()
	dayfuncs[day].(func(string, int))(input, part)
	elapsed := time.Now().Sub(start)
	fmt.Printf("Time: %v\n\n", elapsed)
}

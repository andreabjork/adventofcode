package main

import (
    "adventofcode/m/v2/days"
    "adventofcode/m/v2/util"
    "fmt"
    "os"
    "time"
)

var dayfuncs = map[int]interface{}{
    1: days.Day1,
    2: days.Day2,
    3: days.Day3,
    4: days.Day4,
    5: days.Day5,
    6: days.Day6,
    7: days.Day7,
    8: days.Day8,
	9: days.Day9,
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
        for d := 1; d <= 9; d++ {
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
        test_input = fmt.Sprintf("inputs/day%d_test.txt", day)
        real_input = fmt.Sprintf("inputs/day%d.txt", day)

        if len(os.Args) > 3 && os.Args[3] == "test" {
            runDay(day, part, test_input)
        } else {
            runDay(day, part, test_input)
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
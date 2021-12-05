package main

import (
    "adventofcode/m/v2/days/day1"
    "adventofcode/m/v2/days/day2"
    "adventofcode/m/v2/days/day3"
    "adventofcode/m/v2/days/day4"
    "adventofcode/m/v2/days/day5"
    "fmt"
    "os"
    "time"
)

var dayfuncs = map[string]interface{}{
  "1a": day1.Day1a,
  "1b": day1.Day1b,
  "2a": day2.Day2a,
  "2b": day2.Day2b,
  "3a": day3.Day3a,
  "3b": day3.Day3b,
  "4a": day4.Day4a,
  "4b": day4.Day4b,
  "5a": day5.Day5a,
  "5b": day5.Day5b,
}

func main() {
    if len(os.Args) < 3 {
        fmt.Printf("Run as:\n\tgo run main.go [day][a/b] [input]")
        return
    }

    day := os.Args[1]
    input := fmt.Sprintf("inputs/%s", os.Args[2])
    fmt.Printf("Day: %s, Input: %s\n", day, input)

    start := time.Now()
    dayfuncs[day].(func(string))(input)
    elapsed := time.Now().Sub(start)
    fmt.Printf("\nTime: %v\n", elapsed)
}

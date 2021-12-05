package main

import (
    "adventofcode/m/v2/days"
    "fmt"
    "os"
    "time"
)

var dayfuncs = map[string]interface{}{
  "1a": days.Day1a,
  "1b": days.Day1b,
  "2a": days.Day2a,
  "2b": days.Day2b,
  "3a": days.Day3a,
  "3b": days.Day3b,
  "4a": days.Day4a,
  "4b": days.Undone,
  "5a": days.Day5a,
  "5b": days.Day5b,
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

package main

import (
    "adventofcode/m/v2/days"
    "fmt"
    "os"
)

var dayfuncs = map[string]interface{}{
  "1": days.Undone,
  "2": days.Undone,
  "3": days.Undone,
  "4": days.Day4,
}

func main() {
    if len(os.Args) < 4 {
        fmt.Printf("Run as:\n\tgo run main.go [day] [part] [input]")
        return
    }

    day := os.Args[1]
    part := os.Args[2]
    input := fmt.Sprintf("inputs/%s", os.Args[3])
    fmt.Printf("Day: %s, Part: %s, Input: %s\n", day, part, input)

    dayfuncs[day].(func(string, string))(part, input)
}

package day2

import (
    "adventofcode/m/v2/util"
    "fmt"
    "strconv"
    "strings"
)

func day2a(inputPath string) {

    s := util.LineScanner(inputPath)
    ok := s.Scan()
    line := s.Text()

    course := map[string]int{"forward": 0, "up": 0, "down": 0}
    for ok {
        split := strings.Split(line, " ") 
        if len(split) != 2 {
          break
        }

        key := split[0]
        val, err := strconv.Atoi(split[1])
        if err != nil {
          fmt.Println("Error converting value to int: %+v", split[1])
          panic(err)
        }

        course[key] += val
        ok = s.Scan()
        line = s.Text()
    }
    
    depth := course["down"]-course["up"]
    fmt.Printf("Depth: %d\n", depth)
    fmt.Printf("Forward: %d\n", course["forward"])
    fmt.Printf("Multiple: %d\n", depth*course["forward"])
}

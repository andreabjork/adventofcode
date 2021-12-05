package day2

import (
    "adventofcode/m/v2/util"
    "fmt"
    "strconv"
    "strings"
)

func day2b(inputPath string) {

    s := util.LineScanner(inputPath)
    ok := s.Scan()
    line := s.Text()

    course := map[string]int{"forward": 0, "up": 0, "down": 0, "depth": 0}
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
        
        if key == "forward" {
          course["depth"] += val*(course["down"]-course["up"])
        }

        ok = s.Scan()
        line = s.Text()
    }
    

    fmt.Println(course)
    fmt.Printf("Depth: %d\n", course["depth"])
    fmt.Printf("Forward: %d\n", course["forward"])
    fmt.Printf("Multiple: %d\n", course["depth"]*course["forward"])
}

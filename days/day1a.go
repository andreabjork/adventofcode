package days

import (
    "adventofcode/m/v2/util"
    "fmt"
    "strconv"
)

func Day1a(inputPath string) {

    s := util.LineScanner(inputPath)
    ok := s.Scan()
    line := s.Text()

    // for the first value being counted as
    // an increase
    var (
        count int = -1
        prevVal int = -1
        val int = 0
    )
    for ok {
        val, _ = strconv.Atoi(line)
        if val > prevVal {
            count++
        }

        prevVal = val
        // Read next, if hasNext:
        ok = s.Scan()
        line = s.Text()
    }

    fmt.Printf("Number of increases: %d", count)
}

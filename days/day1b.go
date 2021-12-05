package days

import (
    "adventofcode/m/v2/util"
    "fmt"
    "strconv"
)

func Day1b(inputPath string) {

    s := util.LineScanner(inputPath)
    ok := s.Scan()
    line := s.Text()

    var (
      val int = 0
      numbers []int = make([]int, 0)
    )
    for ok {
        val, _ = strconv.Atoi(line)
        numbers = append(numbers, val)
        ok = s.Scan()
        line = s.Text()
    }
    
    count, err := CountIncreasedWindows(&numbers)
    if err != nil {
      fmt.Println("Error counting number of increases")
      panic(err)
    }
    fmt.Printf("Number of increases: %d", count)
}

func CountIncreasedWindows(numbers *[]int) (int, error) {
    var (
      start int = 0
      end int = 2
      lastSum int = 0
      sum int = 0
      count int = 0
    )

    // Calculate initial sum
    for i := 0; i <= end; i++ {
      sum += (*numbers)[i]
    }

    for end < len(*numbers)-1 {
      lastSum = sum
      end++
      sum += (*numbers)[end]
      sum -= (*numbers)[start]
      start++

      if sum > lastSum {
        count++
      }

    }

    return count, nil
}

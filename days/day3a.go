package main

import (
    "fmt"
  "os"
  "bufio"
  "strconv"
)	

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }

    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Can't open file:", os.Args[1])
        panic(err)
    }

    r := bufio.NewReader(f)
    line, err := Readln(r)
    
    totalbits := len(line)
    oneBitsInColumn := make([]int, totalbits)
    rows := 1 
    for err == nil {
      for i, b := range line {
        oneBitsInColumn [i] += asBit(b)
      }
      line, err = Readln(r)
      rows ++
    }

    gamma := ""
    epsilon := ""
    for _, total := range oneBitsInColumn {
      if total >= rows/2 {
        gamma += "1"
        epsilon += "0"
      } else {
        gamma += "0"
        epsilon += "1"
      }
    }

    fmt.Printf("Gamma: %s\n", gamma)
    fmt.Printf("Epsilon: %s\n", epsilon)

    decGamma, err := strconv.ParseInt(gamma, 2, 64)
    if err != nil {
      fmt.Printf("Cannot convert %s to decimal", gamma)
      panic(err)
    }

    decEpsilon, err := strconv.ParseInt(epsilon, 2, 64)
    if err != nil {
      fmt.Printf("Cannot convert %s to decimal", epsilon)
      panic(err)
    }

    fmt.Printf("Gamma Decimal: %d\n", decGamma)
    fmt.Printf("Epsilon Decimal: %d\n", decEpsilon)
    fmt.Printf("Gamma*Epsilon: %d\n", decGamma*decEpsilon)
}

func asBit(r rune) int {
  if r == '1' {
    return 1
  }

  return 0
}

func Readln(r *bufio.Reader) (string, error) {
  var (isPrefix bool = true
       err error = nil
       line, ln []byte
      )

  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }
  return string(ln),err
}

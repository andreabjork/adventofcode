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

    // Count initialized as -1 to account
    // for the first value being counted as
    // an increase
    var (
        count int = -1
        prevVal int = -1
        val int = 0
    )
    for err == nil {
        val, err = strconv.Atoi(line)
        if val > prevVal {
            count++
        }

        prevVal = val
        // Read next, if hasNext:
        line, err = Readln(r)
    }

    fmt.Printf("Number of increases: %d", count)
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

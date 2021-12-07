package days

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
)

func Day3(inputFile string, part int) {
	if part == 0 {
		day3a(inputFile)
	} else {
		day3b(inputFile)
	}
}

func day3a(inputPath string) {

	s := util.LineScanner(inputPath)
	ok := s.Scan()
	line := s.Text()

	totalbits := len(line)
	oneBitsInColumn := make([]int, totalbits)
	rows := 1
	for ok {
		for i, b := range line {
			oneBitsInColumn [i] += asBit(b)
		}

		ok = s.Scan()
		line = s.Text()
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

	fmt.Printf("Gamma*Epsilon: %d\n", decGamma*decEpsilon)
}

func asBit(r rune) int {
	if r == '1' {
		return 1
	}

	return 0
}

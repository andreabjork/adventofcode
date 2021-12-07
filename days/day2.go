package days

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)

func Day2(inputFile string, part int) {
	if part == 0 {
		day2a(inputFile)
	} else {
		day2b(inputFile)
	}
}

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
	fmt.Printf("Depth*Forward: %d\n", depth*course["forward"])
}

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
	fmt.Printf("Depth*Forward: %d\n", course["depth"]*course["forward"])
}

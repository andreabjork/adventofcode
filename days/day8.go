package days

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strconv"
	"strings"
)


func Day8(inputFile string, part int) {
	if part == 0 {
		day8(inputFile, count1478)
	} else {
		day8(inputFile, totalValue)
	}
}

func day8(inputFile string, d func(s string) int ) {
	ws := util.LineScanner(inputFile)
	line, ok := util.Read(ws)

	var (
		split		[]string
		tenDigits 	[]string
		outcome 	[]string
		sum			= 0
	)
	for ok {
		split = strings.Split(line, " | ")
		tenDigits = strings.Split(split[0], " ")
		outcome = strings.Split(split[1], " ")

		digits := learnCodes(tenDigits)
		reading := decode(outcome, digits)
		sum += d(reading)

		line, ok = util.Read(ws)
	}

	fmt.Println(sum)
}

// Part a) only counts occurrences of 1, 4, 7, 8 in string
func count1478(s string) int {
	return strings.Count(s, "1")+strings.Count(s, "4")+strings.Count(s, "7")+strings.Count(s, "8")
}

// Part b) maps the reading to integer and adds to total sum
func totalValue(s string) int {
	integer, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Couldn't convert string %s to int\n", s)
	}

	return integer
}

// =======
// DECODER
// =======
func decode(outcome []string, digits map[string]*Digit) string {
	// Read rest of the input
	s := ""
	for _, code := range outcome {
		id := validId(digits, code)

		if digit, found := digits[id]; found {
			s += strconv.Itoa((*digit).num)
		} else {
			fmt.Printf("HELP!!!! code %s had typeID %s which matches no digit\n", code, id)
		}
	}

	return s
}
// ===========
// SETUP
// ===========
// A digit has a unique identifier: (length, gcd(4), gcd(7)):
// Use minimal identifiers so we don't need to calculate gcd(4) and gcd(7) for
// all digits
var digitId = []string {
	0: "633", // 0 is formed of 6 edges, 3 in common with 4, 3 in common with 7
	1: "2",
	2: "52",
	3: "533",
	4: "4",
	5: "532",
	6: "632",
	7: "3",
	8: "7",
	9: "64",
}

type Digit struct {
	num 	int
	code	string
}

type UnknownDigit struct {
	code	string
	id 		string // id must be a valid id in the slice above
}

// ===========
// LEARN CODES
// ===========
func learnCodes(words []string) map[string]*Digit {
	digits := make(map[string]*Digit)
	for idx, id := range digitId {
		digits[id] = &Digit{idx, ""}
	}

	unknownDigits := make([]*UnknownDigit, 0)
	// Finds 1, 4, 7, 8, because typeID = len(word) is in the map
	for _, word := range words {
		id := strconv.Itoa(len(word))

		if digit, found := digits[id]; found {
			digit.code = word
		} else {
			unknownDigits = append(unknownDigits, &UnknownDigit{word, id})
		}
	}

	// Finds the rest of the unknown learnCodes
	for _, unknownDigit:= range unknownDigits {
		id := validId(digits, (*unknownDigit).code)
		(*digits[id]).code = (*unknownDigit).code
	}

	return digits
}

// ============
// HELPER FUNCS
// ============
// Tries X, XY, and XYZ until a unique typeID in the map is found
// where X = length of code, Y = gcd(code, 4), Z = gcd(code, 7)
func validId(idmap map[string]*Digit, code string) string {
	id := strconv.Itoa(len(code))
	if _, found := idmap[id]; found {
		return id
	}

	id  += strconv.Itoa(gcd(code, (*idmap[digitId[4]]).code))
	if _, found := idmap[id]; found {
		return id
	}

	id += strconv.Itoa(gcd(code, (*idmap[digitId[7]]).code))
	return id
}

// gcd(a,b) is the number of edges a-digit has in common with b-digit
// Example: gcd(3,4) = 3, because 3-digit has 3 edges in common with the 4-digit
func gcd(a string, b string) int {
	gcd := 0
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			gcd++
		}
	}
	return gcd
}

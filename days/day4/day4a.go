package day4

import (
	"adventofcode/m/v2/util"
	"fmt"
	"strings"
)


func Day4a(inputFile string) {
	ws := util.WordScanner(inputFile)
	line, _ := util.Read(ws)

	// The bingo numbers to mark
	bingoNumbers := strings.Split(line, ",")

	// The bingo number lookup map:
	// map[x] says all sheets and row,col position
	// this number exists on
	bingoLookUp, totalSheets := createBingoLookUpMap(ws)

	// Mark bingoNumbers and update sum of unmarked
	winningNumber, winningSheet, _ := markBingoSheets(bingoNumbers, bingoLookUp, totalSheets, 1)

	fmt.Printf("Bingo! The sum of unmarked numbers on the winning sheet is: %d", winningNumber*winningSheet.sumUnmarked)
}

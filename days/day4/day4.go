package day4

import (
	"adventofcode/m/v2/util"
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func Day4(inputFile string, part int) {
	if part == 0 {
		day4a(inputFile)
	} else {
		day4b(inputFile)
	}
}

func day4a(inputFile string) {
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

	fmt.Printf("First bingo sum: %d\n", winningNumber*winningSheet.sumUnmarked)
}

func day4b(inputFile string) {
	ws := util.WordScanner(inputFile)
	line, _ := util.Read(ws)

	// The bingo numbers to mark
	bingoNumbers := strings.Split(line, ",")

	// The bingo number lookup map:
	// map[x] says all sheets and row,col position
	// this number exists on
	bingoLookUp, totalSheets := createBingoLookUpMap(ws)

	// Mark bingoNumbers and update sum of unmarked
	winningNumber, winningSheet, _ := markBingoSheets(bingoNumbers, bingoLookUp, totalSheets, len(totalSheets))

	fmt.Printf("Last bingo sum: %d\n", winningNumber*winningSheet.sumUnmarked)
}

type BingoPosition struct {
	onSheet	int
	inRow	int
	inCol	int
}

type BingoSheetResult struct {
	xInRow	[]int
	xInCol	[]int
	sumUnmarked	int
	bingoNotFound bool
}

func markBingoSheets(bingoNumbers []string, bingoLookUp map[int][]BingoPosition, marks []BingoSheetResult, stopAt int) (
	int,
	BingoSheetResult,
	error) {

	numSheetsWithBingo := 0
	for _, number := range bingoNumbers {
		num := util.ToInt(number)
		for _, pos := range bingoLookUp[num] {
			if marks[pos.onSheet].bingoNotFound {
				// Update row count
				marks[pos.onSheet].xInRow[pos.inRow]++

				// Update col count
				marks[pos.onSheet].xInCol[pos.inCol]++

				// Update sum of unmarked
				marks[pos.onSheet].sumUnmarked -= num

				if marks[pos.onSheet].xInRow[pos.inRow] == 5 || marks[pos.onSheet].xInCol[pos.inCol] == 5 {
					numSheetsWithBingo++
					marks[pos.onSheet].bingoNotFound = false
					if numSheetsWithBingo == stopAt {
						return num, marks[pos.onSheet], nil
					}
				}
			}
		}
	}

	return -1, BingoSheetResult{}, errors.New("No bingo found.")
}

// x["22"] = [
//   { 3, 1, 2 },
//   { 1, 1, 1 } ]
// means the number "22" is on bingo sheet 3 at row 2, column 3,
// and bingo sheet 1 at row 1, column 1.
func createBingoLookUpMap(ws *bufio.Scanner) (map[int][]BingoPosition, []BingoSheetResult) {
	sheets := make([]BingoSheetResult, 0)

	bingoLookUp := make(map[int][]BingoPosition)
	var (
		bingoNumber int = 0
		sheetNumber int = 0
		row         int = 0
		col         int = 0
		count       int = 0
	)
	word, ok := util.Read(ws)
	sum := 0
	for ok {
		bingoNumber = util.ToInt(word)
		sum += bingoNumber
		bingoLookUp[bingoNumber] = append(bingoLookUp[bingoNumber], BingoPosition{
			onSheet: sheetNumber,
			inRow:   row,
			inCol:   col,
		})

		// increments
		count++
		col++
		if count%5 == 0 {
			row++
			col = 0
		}
		if count == 25 {
			sheets = append(sheets, BingoSheetResult{
				xInRow:      make([]int, 5),
				xInCol:      make([]int, 5),
				sumUnmarked: sum,
				bingoNotFound: true,
			})
			sheetNumber++
			count = 0
			row = 0
			col = 0
			sum = 0
		}

		word, ok = util.Read(ws)
	}

	return bingoLookUp, sheets
}
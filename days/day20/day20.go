package day20

import (
	"adventofcode/m/v2/util"
	"fmt"
)

func Day20(inputFile string, part int) {
	// Works for n = even, n = odd with parity 0 will lead to infinite
	// lit pixels.
	img, algo := MakeImage(inputFile)
	if part == 0 {
		fmt.Printf("Solution: %d\n", img.enhanceTimes(algo, 2))
	} else {
		fmt.Printf("Solution: %d\n", img.enhanceTimes(algo, 50))
	}
}

func MakeImage(inputFile string) (*Image, []bool) {
	ls := util.LineScanner(inputFile)

	// Process algorithm
	algo := make([]bool, 512)
	line, _ := util.Read(ls)
	runes := []rune(line)
	for i, r := range runes {
		algo[i] = r == '#'
	}
	// read empty line
	_,_ = util.Read(ls)

	// read image, coordinate system:
	// (0,0) (0,1) (0,2) ...
	// (1,0) (1,1) (1,2) ...
	pixels := map[int]map[int]bool{}
	col := 0
	row := 0
	for line, ok := util.Read(ls); ok; line, ok = util.Read(ls) {
		runes = []rune(line)
		if len(runes) > col {
			col = len(runes)
		}
		pixels[row] = make(map[int]bool, len(runes))
		for col, r := range runes {
			pixels[row][col] = r == '#'
		}
		row++
	}

	return &Image{pixels, 0, col, row}, algo
}

type Image struct {
	// Coordinate system originating in top left
	pixels 	map[int]map[int]bool
	parity  int
	width 	int
	height	int
}

func (img *Image) print() {
	fmt.Printf("Printing %d by %d image\n", img.width, img.height)
	for i := 0; i < img.height; i++ {
		for j := 0; j < img.width; j++ {
			if img.pixels[i][j] {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println("")
	}
}

func (img *Image) padding() *Image {
	m, n := img.width, img.height
	paddedPixels := make(map[int]map[int]bool, m+2)

	for i := 0; i < m+2; i++ {
		paddedPixels[i] = make(map[int]bool, n+2)
		for j := 0; j < n+2; j++ {
			paddedPixels[i][j] = img.pixel(i-1,j-1) == 1
		}
	}

	return &Image{paddedPixels, img.parity, m+2, n+2}
}

func (img *Image) pixel(i int, j int) int {
	if i >= 0 && i <= img.width && j >= 0 && j <= img.height {
		if val, ok := img.pixels[i][j]; val && ok{
			return 1
		} else if ok {
			return 0
		}
	}
	// images are infinite with either all lit or all unlit pixels
	return img.parity
}

func (img *Image) enhanceTimes(algo []bool, n int) int {
	var litPixels int
	enhancedImg := img
	for i := 0; i < n; i++ {
		enhancedImg = enhancedImg.padding()
		enhancedImg, litPixels = enhancedImg.enhance(algo)
	}
	return litPixels
}

func (img *Image) enhance(algo []bool) (*Image, int) {
	litPixels := 0
	enhancedImg := newImage(img.width, img.height)
	for i := 0; i < img.height; i++ {
		for j := 0; j < img.width; j++ {
			enhancedImg.pixels[i][j] = img.enhancedPixel(i,j, algo)
			if enhancedImg.pixels[i][j] {
				litPixels++
			}
		}
	}

	if img.parity == 0 {
		if algo[0] {
			enhancedImg.parity = 1
		} else {
			enhancedImg.parity = 0
		}
	} else {
		if algo[len(algo)-1] {
			enhancedImg.parity = 1
		} else {
			enhancedImg.parity = 0
		}
	}

	return enhancedImg, litPixels
}

func (img *Image) enhancedPixel(m int, n int, algo []bool) bool {
	return algo[img.pixelDecimal(m,n)]
}

func (img *Image) pixelDecimal(row int, col int) int {
	// 3x3 window is binary string b1 b2 ... b9
	// ----------------------------------------
	// (0,0) (0,1) (0,2) ...
	// (1,0) (1,1) (1,2) ...
	// (2,0) (2,1) (2,2) ...
	// is processed in the order of b9 b8 ... b1
	// that is (2,2) (2,1) (2,0) (1,2) ...

	b := 0
	decimal := 0
	for i := row+1; i >= row-1; i-- {
		for j := col+1; j >= col-1; j-- {
			decimal += img.pixel(i,j)*pow2(b)
			b++
		}
	}
	return decimal
}

func newImage(width int, height int) *Image {
	pixels := make(map[int]map[int]bool, width)
	for i := 0; i < width; i++ {
		pixels[i] = make(map[int]bool, height)
	}

	return &Image{pixels, 0, width, height}
}

func pow2(n int) int {
	if n == 0 {
		return 1
	}

	return 2*pow2(n-1)
}
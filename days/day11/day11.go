package day11

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
	"strconv"
)

func Day11(inputFile string, part int) {
	if part == 0 {
		day11a(inputFile)
	} else {
		day11b(inputFile)
	}
}

func day11a(inputFile string) {
	ls := util.LineScanner(inputFile)

	c := Colony{make([][]*Octopus, 0), make([]*Octopus, 0), 0}
	c.populateColony(ls)
	for i := 1; i <= 100; i++ {
		c.timePasses()
		for _, octopus := range c.flashing {
			octopus.flash(&c)
		}
	}
	fmt.Printf("Flashes: %d\n", c.flashes)
}

func day11b(inputFile string) {
	ls := util.LineScanner(inputFile)

	c := Colony{make([][]*Octopus, 0), make([]*Octopus, 0), 0}
	c.populateColony(ls)

	// They're synchronized the first time # flashes == # octopi
	numOctopi := len(c.octopi)*len(c.octopi[0])
	synchronized := -1
	for i := 1; synchronized == -1; i++ {
		flashesBefore := c.flashes
		c.timePasses()
		for _, octopus := range c.flashing {
			octopus.flash(&c)
		}

		if c.flashes-flashesBefore == numOctopi {
			synchronized = i
			break
		}
	}

	fmt.Printf("Synchronized at step: %d\n", synchronized)
}

// =============
// OCTOPI COLONY
// =============
const FLASH_ENERGY = 10

type Colony struct {
	octopi    [][]*Octopus
	flashing  []*Octopus
	flashes   int
}

func (c *Colony) populateColony(ls *bufio.Scanner) {
	line, ok := util.Read(ls)

	// Read in all the octopi
	for i := 0; ok; i++ {
		(*c).octopi = append((*c).octopi, make([]*Octopus, 0))
		energies := []rune(line)

		for j, erune := range energies {
			e, _ := strconv.Atoi(string(erune))
			o := &Octopus{e, i, j, nil, false}
			(*c).octopi[i] = append((*c).octopi[i], o)
			if (*o).energy == FLASH_ENERGY {
				(*c).flashing = append((*c).flashing, o)
			}
		}

		line, ok = util.Read(ls)
	}
}

func (c *Colony) of(x int, y int) (bool, *Octopus){
	if x < 0 || x >= len((*c).octopi) || y < 0 || y >= len((*c).octopi) {
		return false, nil
	}
	return true, (*c).octopi[x][y]
}

func (c *Colony) timePasses() {
	(*c).flashing = make([]*Octopus, 0)
	for i := 0; i < len((*c).octopi); i++{
		for j := 0; j < len((*c).octopi[i]); j++ {
			o := (*c).octopi[i][j]
			(*o).flashed = false
			(*o).energy++
			if (*o).energy == FLASH_ENERGY {
				(*c).flashing = append((*c).flashing, o)
			} else if (*o).energy > FLASH_ENERGY {
				(*o).energy = 0
			}
		}
	}
}

func (c *Colony) show(text string) {

	fmt.Printf("========== %s ==========\n", text)
	for i := 0; i < len((*c).octopi); i++ {
		for j := 0; j < len((*c).octopi[i]); j++ {
			fmt.Printf(" %d ", (*(*c).octopi[i][j]).energy)
		}
		fmt.Println("")
	}
}

// =======
// OCTOPUS
// =======
type Octopus struct {
	energy 	int
	x	int
	y 	int
	neighbours	[]*Octopus
	flashed	bool
}

func (o *Octopus) flash(c *Colony) {
	(*o).flashed = true
	(*o).energy = 0
	(*c).flashes++
	if (*o).neighbours == nil {
		(*o).findNeighbours(c)
	}
	for _, n := range (*o).neighbours {
		if !(*n).flashed {
			(*n).energy++
			if (*n).energy == FLASH_ENERGY {
				n.flash(c)
			}
		}
	}
}

func (o *Octopus) findNeighbours(colony *Colony) {
	neighbours := make([]*Octopus, 0)
	for X := (*o).x-1; X <= (*o).x+1; X++ {
		for Y := (*o).y-1; Y <= (*o).y+1; Y++ {
			if exists, neighbour := colony.of(X,Y); exists && neighbour != o {
				neighbours = append(neighbours, neighbour)
			}
		}
	}

	(*o).neighbours = neighbours
}

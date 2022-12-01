package day19

import (
	"adventofcode/m/v2/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var MAGIC_NUM int
var rotations map[int]func(c *Coord) (*Coord)

func Day19(inputFile string, part int) {
	MAGIC_NUM = 11
	if part == 0 {
		fmt.Printf("Solution: %d\n", Solve(inputFile))
	} else {
		fmt.Printf("not impl")
	}
}

func Solve(inputFile string) int {

	rotations = map[int]func(c *Coord) *Coord {
		// 0 degree clockwise is x,y -> x,y
		1: func(c *Coord) *Coord { return &Coord{c.x, c.y, c.z} },
		2:	func(c *Coord) *Coord { return &Coord{c.z,c.x,c.y }},
		3:	func(c *Coord) *Coord { return &Coord{c.y,c.z,c.x }},
		// 90 degree clockwise is x,y -> y,-x
		4:	func(c *Coord) *Coord { return &Coord{c.y,-c.x,c.z }},
		5:	func(c *Coord) *Coord { return &Coord{c.z,c.y,-c.x }},
		6:	func(c *Coord) *Coord { return &Coord{-c.x,c.z,c.y }},
		// 180 degree clockwise (upside down) is x,y -> -x,-y
		7:	func(c *Coord) *Coord { return &Coord{-c.x,-c.y,c.z }},
		8:	func(c *Coord) *Coord { return &Coord{c.z,-c.x,-c.y }},
		9:	func(c *Coord) *Coord { return &Coord{-c.y,c.z,-c.x }},
		// 270 degree clockwise is -x,y -> -y,x
		10:	func(c *Coord) *Coord { return &Coord{-c.y,c.x,c.z }},
		11:	func(c *Coord) *Coord { return &Coord{c.z,-c.y,c.x }},
		12:	func(c *Coord) *Coord { return &Coord{c.x,c.z,-c.y }},
		// 0 clockwise, flipped
		13:	func(c *Coord) *Coord { return &Coord{c.y,c.x,-c.z }},
		14:	func(c *Coord) *Coord { return &Coord{-c.z,c.y,c.x} },
		15:	func(c *Coord) *Coord { return &Coord{c.x,-c.z,c.y} },
		// 90 degree clockwise, flipped, x,y,z -> y,-x,-z
		16:	func(c *Coord) *Coord { return &Coord{-c.x,c.y,-c.z} },
		17:	func(c *Coord) *Coord { return &Coord{-c.z,-c.x,c.y }},
		18:	func(c *Coord) *Coord { return &Coord{c.y,-c.z,-c.x }},
		// 180 degree clockwise (upside down) is x,y -> -x,-y
		19:	func(c *Coord) *Coord { return &Coord{-c.y,-c.x,-c.z }},
		20:	func(c *Coord) *Coord { return &Coord{-c.z,-c.y,-c.x }},
		21:	func(c *Coord) *Coord { return &Coord{-c.x,-c.z,-c.y }},
		// 270 degree clockwise is -x,y -> -y,x
		22:	func(c *Coord) *Coord { return &Coord{c.x,-c.y,-c.z }},
		23:	func(c *Coord) *Coord { return &Coord{-c.z,c.x,-c.y }},
		24:	func(c *Coord) *Coord { return &Coord{-c.y,-c.z,c.x }},

	}

	beaconMap := &Map{area: map[int]map[int]map[int]bool{}, beacons: []*Coord{}}
	scanners := ReadPositions(inputFile)
	orientedScanners := []*Scanner{scanners[0]}

	// Initializing map based on origin in S_0
	scanners[0].pos = &Coord{0,0,0}
	scanners[0].rotation = 1
	beaconMap.addAll(scanners[0])

	for len(orientedScanners) < len(scanners) {
		for _, t := range scanners {
			// t has not been mapped yet
			if t.pos == nil {
				// find pairing within oriented scanners
				for _, s := range orientedScanners {
					if s.pairsWith(t) {
						orientedScanners = append(orientedScanners, t)
						beaconMap.addAll(t)
						break
					}
				}
			}
		}

		// Find next scanner pair
	}

	return len(beaconMap.beacons)
}



func ReadPositions(inputFile string) []*Scanner {
	ls := util.LineScanner(inputFile)

	scanners := []*Scanner{}
	line, ok := util.Read(ls)
	re := regexp.MustCompile("--- scanner [0-9]+ ---")
	i := -1
	for ok {
		next := re.MatchString(line)
		if next {
			i++
			//fmt.Printf("--- Scanner %d ---\n", i)
			scanners = append(scanners, NewScanner(i))
			line, ok = util.Read(ls)
			continue
		} else {
			coords := strings.Split(line, ",")
			fmt.Printf("%s\n", line)
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			z, _ := strconv.Atoi(coords[2])
			scanners[i].beacons = append(scanners[i].beacons, &Coord{x, y, z})

			line, ok = util.Read(ls)
			fmt.Println(line)
			if line == "" {
				line, ok = util.Read(ls)
			}
		}
	}

	for _, scanner := range scanners {
		scanner.calcDistances()
	}

	return scanners
}

type Tuple struct {
	i 	int
	j 	int
}

type Coord struct {
	x	int
	y	int
	z	int
}

func (c *Coord) add(a *Coord) *Coord {
	return &Coord{c.x+a.x, c.y+a.y, c.z+a.z}
}

func (c *Coord) subtract(a *Coord) *Coord {
	return &Coord{c.x-a.x, c.y-a.y, c.z-a.z}
}

func (c *Coord) print() string {
	return fmt.Sprintf("(%d,%d,%d)", c.x, c.y, c.z)
}

type Map struct {
	area map[int]map[int]map[int]bool
	beacons []*Coord
}

func (m *Map) addAll(t *Scanner){
	for _,b := range t.beacons {
		// ST + (TB)v where S is the origin, so ST is T's true position
		beacon := t.pos.add(rotations[t.rotation](b))
		if ! m.lookup(beacon) {
			m.add(beacon)
		}
	}
}

func (m *Map) lookup(b *Coord) bool {
	_, ok := m.area[b.x][b.y][b.z]
	return ok
}

func (m *Map) add(b *Coord) {
	if _, ok := m.area[b.x]; ok {
		if _, ok := m.area[b.x][b.y]; ok {
			m.area[b.x][b.y][b.z] = true
		} else {
			m.area[b.x][b.y] = map[int]bool{}
			m.area[b.x][b.y][b.z] = true
		}
	} else {
		m.area[b.x] = map[int]map[int]bool{}
		m.area[b.x][b.y] = map[int]bool{}
		m.area[b.x][b.y][b.z] = true
	}

	m.beacons = append(m.beacons, &Coord{b.x, b.y, b.z})
}

type Scanner struct {
	pos			*Coord  // nil to start, need to find these
	rotation	int // 0 to start, need to find these
	id 			int
	beacons		[]*Coord
	dsq 		[][]int
}

func NewScanner(id int) *Scanner {
	return &Scanner{nil, 0, id, make([]*Coord, 0), nil}
}

func (s *Scanner) calcDistances() {
	var dist int
	s.dsq = make([][]int, 0)
	for i := 0; i < len(s.beacons); i++ {
		s.dsq = append(s.dsq, make([]int, 0))
		for j := 0; j < len(s.beacons); j++ {
			a := s.beacons[i]
			b := s.beacons[j]
			dist = (a.x - b.x)*(a.x - b.x) + (a.y - b.y)*(a.y - b.y) + (a.z - b.z)*(a.z - b.z)
			s.dsq[i] = append(s.dsq[i], dist)
		}
	}
}

func (s *Scanner) pairsWith(t *Scanner) bool {
	// Find 2 beacons, b1 and c1, such that the union of their distances to the other beacons around
	// them contains at least 12 beacons: Then maybe SB == BT and this will be verified later.
	paired := true
	for _, sdsq := range s.dsq {
		for _, tdsq := range t.dsq {
			maybePair, union := Union(sdsq, tdsq)
			if maybePair {
				paired = Validate(union, s, t)
				if paired {
					return true
				}
			}
		}
	}

	return false
}

func Validate(union map[int][]*Tuple, s *Scanner, t *Scanner) bool {
	// Calculate t.pos based on all beacons in (the potential) union
	// If we find 12 equal t.pos, then those 12 form the union and that is the true position of t.
	// Remember
	// 	 ST = SB + (BT)v, for every beacon B in their union where BT has been rotated to match ST.
	validated := false
	for rot_id := 1; rot_id <= 24; rot_id++ {
		rotation := rotations[rot_id]
		t_pos := map[int]map[int]map[int]int{}
		//fmt.Printf("\n\nFor rotation s,t: %+v, %+v\n", s.rotation, rot_id)
		for _, tuples := range union {
			for _, pair := range tuples {
				// Condition: Scanner s' data has been been centered in O and oriented to match O.
				// ST = (SB)v + (BT)w = (SB)v - (TB)w
				b := s.beacons[pair.i]
				c := t.beacons[pair.j]
				sb := rotations[s.rotation](b)
				tb := rotation(c)
				ST := sb.subtract(tb)
				//fmt.Printf("Beacon %s rotated to %s: ", c.print(), tb.print())
				//fmt.Printf("S%d: %s \n", t.id, ST.print())

				if _, ok := t_pos[ST.x]; ok {
					if _, ok := t_pos[ST.x][ST.y]; !ok {
						t_pos[ST.x][ST.y] = map[int]int{}
					}
				} else {
					t_pos[ST.x] = map[int]map[int]int{}
					t_pos[ST.x][ST.y] = map[int]int{}
				}

				t_pos[ST.x][ST.y][ST.z]++
				if t_pos[ST.x][ST.y][ST.z] >= MAGIC_NUM {
					validated = true
					//t.pos = &Coord{ST.x, ST.y, ST.z}
					t.pos = ST.add(s.pos)
					t.rotation = rot_id
					fmt.Printf("\nValidated scanner S%d\n", t.id)
					fmt.Printf("Scanner position:  (%d,%d,%d)\n", t.pos.x, t.pos.y, t.pos.z )
					fmt.Printf("Scanner rotation: %+v\n", t.rotation)

					break
				}
			}
			if validated {
				break
			}
		}
		if validated {
			break
		}
	}

	return validated
}

func Union(v []int, w []int) (bool, map[int][]*Tuple) {
	size := 0
	union :=map[int][]*Tuple{}
	for i, ival := range v {
		for j := 0; j < len(w); j++ {
			if ival == w[j] {
				if _, ok := union[ival]; !ok {
					union[ival] = make([]*Tuple, 0)
				}

				union[ival] = append(union[ival], &Tuple{i, j})
				size++
			}
		}
	}

	return size >= MAGIC_NUM, union
}


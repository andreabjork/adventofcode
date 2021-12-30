package day15

import (
	"adventofcode/m/v2/util"
	"bufio"
	"fmt"
	"strconv"
)

func Day15(inputFile string, part int) {
	ls := util.LineScanner(inputFile)
	cave := CCave{make([][]*CNode, 0), make([]*CNode, 0)}
	cave.mapCave(ls)
	if part == 1 {
		// Make cave 5x5, with weights wrapping around 1,2,...,9,1,2...,9,1...
		cave.resizeCave()
	}
	_, start := cave.of(0,0)
	start.shortest = 0
	cave.toConsider = append(cave.toConsider, start)
	cave.markShortest()
	fmt.Printf("Shortest path to last: %d \n", cave.nodes[len(cave.nodes)-1][len(cave.nodes[0])-1].shortest)
}

type CCave struct {
	nodes 		[][]*CNode
	toConsider	[]*CNode
}

func (c *CCave) of(x int, y int) (bool, *CNode){
	if x < 0 || x >= len(c.nodes) || y < 0 || y >= len(c.nodes[0]) {
		return false, nil
	}
	return true, c.nodes[x][y]
}

func (c *CCave) mapCave(ls *bufio.Scanner) {
	line, ok := util.Read(ls)

	row := 0
	MAX_INT := int(^uint(0) >> 1)
	for ok {
		c.nodes = append(c.nodes, make([]*CNode, 0))
		runes := []rune(line)
		for col, r := range runes {
			weight, _ := strconv.Atoi(string(r))
			cnode := CNode{weight, MAX_INT,row, col, make([]*CNode, 0)}
			c.nodes[row] = append(c.nodes[row], &cnode)
		}

		row++
		line, ok = util.Read(ls)
	}
}

func (c *CCave) resizeCave() {
	MAX_INT := int(^uint(0) >> 1)
	N := len(c.nodes)
	M := len(c.nodes[0])
	nodes := make([][]*CNode, N*5)
	for i := 0; i < N*5; i++ {
		nodes[i] = make([]*CNode, M*5)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++{

					weight := (c.nodes[i][j].weight+k+l-1)%9+1
					cnode := CNode{weight, MAX_INT,i+k*N, j+l*M, make([]*CNode, 0)}
					nodes[i+k*N][j+l*M] = &cnode
				}
			}
		}
	}

	c.nodes = nodes
}

func (c *CCave) markShortest() {
	idx := 0
	for idx < len(c.toConsider)  {
		node := c.toConsider[idx]
		if len(node.neighbours) == 0 {
			node.findNeighbours(c)
		}
		for _, n := range node.neighbours {
			if node.shortest+n.weight < n.shortest {
				n.shortest = node.shortest+n.weight
				c.toConsider = append(c.toConsider, n)
			}
		}
		idx++
	}
}

type CNode struct {
	weight 		int
	shortest	int
	x 			int
	y 			int
	neighbours	[]*CNode
}

func (cn *CNode) findNeighbours(cave *CCave) {
	// Right, left, bottom, top; diagonals not considered
	if exists, neighbour := cave.of(cn.x+1,cn.y); exists && neighbour != cn {
		cn.neighbours = append(cn.neighbours, neighbour)
	}
	if exists, neighbour := cave.of(cn.x-1,cn.y); exists && neighbour != cn {
		cn.neighbours = append(cn.neighbours, neighbour)
	}
	if exists, neighbour := cave.of(cn.x,cn.y+1); exists && neighbour != cn {
		cn.neighbours = append(cn.neighbours, neighbour)
	}
	if exists, neighbour := cave.of(cn.x,cn.y-1); exists && neighbour != cn {
		cn.neighbours = append(cn.neighbours, neighbour)
	}
}

package days

import (
  "adventofcode/m/v2/util"
  "bufio"
  "fmt"
  "strconv"
)

func day3b(inputPath string) {
    s := util.LineScanner(inputPath)
    root := createTree(s)

    oxygen_generator_rating := root.traverseHeaviest("")
    co2_scrubber_rating := root.traverseLightest("")

    decOGR, err := strconv.ParseInt(oxygen_generator_rating, 2, 64)
    if err != nil {
      fmt.Println("Error converting oxygen generator rating %s to binary", oxygen_generator_rating)
    }

    decCO2SR, err := strconv.ParseInt(co2_scrubber_rating, 2, 64)
    if err != nil {
      fmt.Println("Error converting co2 scrubber rating %s to binary", co2_scrubber_rating)
    }

    strconv.ParseInt(oxygen_generator_rating, 2, 64)
    fmt.Printf("OGR*CO2SR: %d\n", decOGR*decCO2SR)
}


type Node struct {
  weight  int // The weight of a node is the size of the node's subtree.
  left    *Node
  right   *Node
}

// Create a tree. Whenever we add a binary number to the tree,
// traverse one edge to the left if the next digit is 0, otherwise
// one edge to the right. Update the weight of each node you traverse.
//
func createTree(s *bufio.Scanner) (*Node) {
  ok := s.Scan()
  binary := s.Text()

  rootNode := &Node{weight: 0, left: nil, right: nil}
  node := rootNode

  // Add every binary as a leaf
  for ok {
    // Traverse left or right depending on bit,
    // and update weights
    for _, bit := range binary {
      if bit == '0' {
        node = node.moveLeft() // adds +1 to left weight
      } else {
        node = node.moveRight() // adds +1 to right weight
      }
    }

    node = rootNode
    ok = s.Scan()
    binary = s.Text()
  }

  return rootNode
}

func (n *Node) moveLeft() (*Node) {
  (*n).weight++
  if (*n).left == nil {
    (*n).left = &Node{weight: 0, left: nil, right: nil}
  }

  return (*n).left
}

func (n *Node) moveRight() (*Node) {
  (*n).weight++
  if (*n).right == nil {
    (*n).right = &Node{weight: 0, left: nil, right: nil}
  }

  return (*n).right
}

func (n Node) traverseLightest(path string) (string) {
  if n.left == nil && n.right == nil {
    return path
  }
  
  next := n.right
  if n.left == nil && n.right != nil {
    path += "1"
  } else if n.left != nil && n.right == nil {
    path += "0"
    next = n.left
  } else if n.left.weight <= n.right.weight {
    path += "0"
    next = n.left
  } else {
    path += "1"
  }

  return (*next).traverseLightest(path)
}

func (n Node) traverseHeaviest(path string) (string) {
  if n.left == nil && n.right == nil {
    return path
  }
  
  next := n.right
  if n.left == nil && n.right != nil {
    path += "1"
  } else if n.left != nil && n.right == nil {
    path += "0"
    next = n.left
  } else if n.right.weight >= n.left.weight {
    path += "1"
  } else {
    path += "0"
    next = n.left
  }

  return (*next).traverseHeaviest(path)
}

package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)	

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }

    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Can't open file:", os.Args[1])
        panic(err)
    }

    r := bufio.NewReader(f)

    root := createTree(r)
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
    fmt.Printf("Oxygen Generator Rating: %s = %d\n", oxygen_generator_rating, decOGR)
    fmt.Printf("CO2 Scrubber Rating: %s = %d\n", co2_scrubber_rating, decCO2SR)
    fmt.Printf("Multiplied value: %d\n", decOGR*decCO2SR)
}


type Node struct {
  weight  int // The weight of a node is the size of the node's subtree.
  left    *Node
  right   *Node
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


// Create a tree. Whenever we add a binary number to the tree,
// traverse one edge to the left if the next digit is 0, otherwise
// one edge to the right. Update the weight of each node you traverse.
//
func createTree(r *bufio.Reader) (*Node) {
  binary, err := Readln(r)
  rootNode := &Node{weight: 0, left: nil, right: nil}
  node := rootNode

  // Add every binary as a leaf
  for err == nil {
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
    binary, err = Readln(r)
  }
  
  return rootNode
}

func Readln(r *bufio.Reader) (string, error) {
  var (isPrefix bool = true
       err error = nil
       line, ln []byte
      )

  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }
  return string(ln),err
}

package day08

import "fmt"

// Node represents a node in a map.
type Node struct {
	Label string
	Left  string
	Right string
}

// NewNode creates a new node.
func NewNode(input string) *Node {
	return &Node{
		Label: input[0:3],
		Left:  input[7:10],
		Right: input[12:15],
	}
}

// Map represents a map of nodes.
type Map struct {
	Nodes map[string]*Node
}

// NewMap creates a new map.
func NewMap() *Map {
	return &Map{
		Nodes: make(map[string]*Node),
	}
}

// AddNode adds a node to the map.
func (m *Map) AddNode(input string) {
	node := NewNode(input)
	m.Nodes[node.Label] = node
}

// NumStepsToDestination returns the number of steps required to reach the destination.
func (m *Map) NumStepsToDestination(directions []rune) int {
	result := 0
	i := 0
	node := m.Nodes["AAA"]
	for {
		if directions[i] == 'L' {
			node = m.Nodes[node.Left]
		} else if directions[i] == 'R' {
			node = m.Nodes[node.Right]
		} else {
			panic(fmt.Sprintf("invalid direction: %c", directions[i]))
		}
		result += 1
		i = (i + 1) % len(directions)
		if node.Label == "ZZZ" {
			return result
		}
	}
}

// NumStepsToDestination returns the number of steps required to reach the destination.
func NumStepsToDestination(input []string) int {
	directions := []rune(input[0])

	m := NewMap()
	for _, line := range input[2:] {
		m.AddNode(line)
	}

	return m.NumStepsToDestination(directions)
}

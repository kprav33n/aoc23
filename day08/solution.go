package day08

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

// NumStepsToPredicate returns the number of steps required to reach the predicate.
func (m *Map) NumStepsToPredicate(start string, directions []rune, predicate func(*Node) bool) int {
	i, result := 0, 0
	for node := m.Nodes[start]; predicate(node) == false; {
		if directions[i] == 'L' {
			node = m.Nodes[node.Left]
		} else {
			node = m.Nodes[node.Right]
		}
		result += 1
		i = (i + 1) % len(directions)
	}
	return result
}

// LCM returns the least common multiple of the given numbers.
func LCM(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = result * number / GCD(result, number)
	}
	return result
}

// GCD returns the greatest common divisor of the given numbers.
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// NumStepsToDestination returns the number of steps required to reach the destination.
func NumStepsToDestination(input []string) int {
	directions := []rune(input[0])

	m := NewMap()
	for _, line := range input[2:] {
		m.AddNode(line)
	}

	return m.NumStepsToPredicate("AAA", directions, func(n *Node) bool {
		return n.Label == "ZZZ"
	})
}

// NumStepsToDestinationForGhost returns the number of steps required to reach
// the destination for the ghost.
func NumStepsToDestinationForGhost(input []string) int {
	directions := []rune(input[0])

	m := NewMap()
	for _, line := range input[2:] {
		m.AddNode(line)
	}

	var numStepsList []int
	for _, node := range m.Nodes {
		if node.Label[2:] == "A" {
			n := m.NumStepsToPredicate(node.Label, directions, func(n *Node) bool {
				return n.Label[2:] == "Z"
			})
			numStepsList = append(numStepsList, n)
		}
	}
	return LCM(numStepsList)
}

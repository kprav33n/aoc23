package day10

// Point represents a point.
type Point struct {
	X int
	Y int
}

// Add returns the sum of two points.
func (p *Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Difference returns the difference of two points.
func (p *Point) Difference(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// Pipe represents a pipe.
type Pipe struct {
	From Point
	To   Point
	Char string
}

var pipeMapping = map[rune]Pipe{
	'|': {Point{0, -1}, Point{0, 1}, "|"},
	'-': {Point{-1, 0}, Point{1, 0}, "-"},
	'L': {Point{0, -1}, Point{1, 0}, "L"},
	'J': {Point{0, -1}, Point{-1, 0}, "J"},
	'7': {Point{-1, 0}, Point{0, 1}, "7"},
	'F': {Point{0, 1}, Point{1, 0}, "F"},
	'.': {Point{0, 0}, Point{0, 0}, "."},
	'S': {Point{0, 0}, Point{0, 0}, "S"},
}

// NewPipe returns a new pipe.
func NewPipe(r rune) Pipe {
	return pipeMapping[r]
}

// Sketch represents a sketch.
type Sketch struct {
	Start Point
	Pipes [][]Pipe
}

// NewSketch returns a new sketch.
func NewSketch(input []string) Sketch {
	sketch := Sketch{}
	var start Point
	for _, line := range input {
		var row []Pipe
		for _, r := range line {
			if r == 'S' {
				start = Point{len(row), len(sketch.Pipes)}
			}
			row = append(row, NewPipe(r))
		}
		sketch.Pipes = append(sketch.Pipes, row)
	}
	sketch.Start = start
	return sketch
}

// ConnectedToStart returns the pipes that are connected to start.
func (sketch *Sketch) ConnectedToStart() ([]Point, []Point) {
	candidates := []Point{}
	diffs := []Point{}
	minX, minY := 0, 0
	maxX, maxY := len(sketch.Pipes[0]), len(sketch.Pipes)
	offsets := []Point{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	for _, offset := range offsets {
		p := sketch.Start.Add(offset)
		if p.X < minX || p.X >= maxX || p.Y < minY || p.Y >= maxY {
			continue
		}
		diff := sketch.Start.Difference(p)
		pipe := sketch.Pipes[p.Y][p.X]
		if pipe.From == diff || pipe.To == diff {
			candidates = append(candidates, p)
			diffs = append(diffs, diff)
		}
	}
	return candidates, diffs
}

// BackToStart returns the number of steps back to start.
func (sketch *Sketch) BackToStart(p Point, diff Point, acc int) int {
	if p == sketch.Start {
		return acc
	}
	pipe := sketch.Pipes[p.Y][p.X]
	var next Point
	if diff == pipe.From {
		next = p.Add(pipe.To)
	} else {
		next = p.Add(pipe.From)
	}
	return sketch.BackToStart(next, p.Difference(next), acc+1)
}

// StepsToFarthestPoint returns the number of steps to the farthest point.
func StepsToFarthestPoint(input []string) int {
	sketch := NewSketch(input)
	neighbors, diffs := sketch.ConnectedToStart()
	n := sketch.BackToStart(neighbors[0], diffs[0], 1)
	return n / 2
}

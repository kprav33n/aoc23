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

// StartPipe returns the pipe at start.
func PipeForDiff(x, y Point) Pipe {
	for _, pipe := range pipeMapping {
		if (pipe.From == x && pipe.To == y) || (pipe.From == y && pipe.To == x) {
			return pipe
		}
	}
	return pipeMapping['.']
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

	startPipe := PipeForDiff(candidates[0].Difference(sketch.Start), candidates[1].Difference(sketch.Start))
	sketch.Pipes[sketch.Start.Y][sketch.Start.X] = startPipe

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

// PathToStart returns the path back to start.
func (sketch *Sketch) PathToStart(p Point, diff Point, acc []Point) []Point {
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
	return sketch.PathToStart(next, p.Difference(next), append(acc, next))
}

// StepsToFarthestPoint returns the number of steps to the farthest point.
func StepsToFarthestPoint(input []string) int {
	sketch := NewSketch(input)
	neighbors, diffs := sketch.ConnectedToStart()
	n := sketch.BackToStart(neighbors[0], diffs[0], 1)
	return n / 2
}

// NumEnclosedTiles returns the number of enclosed tiles.
func NumEnclosedTiles(input []string) int {
	sketch := NewSketch(input)
	neighbors, diffs := sketch.ConnectedToStart()
	acc := []Point{neighbors[0]}
	path := sketch.PathToStart(neighbors[0], diffs[0], acc)
	pathMap := make(map[Point]struct{})
	for _, p := range path {
		pathMap[p] = struct{}{}
	}

	enclosed := 0
	isInside := false
	for y, row := range sketch.Pipes {
		for x, pipe := range row {
			if _, ok := pathMap[Point{x, y}]; ok {
				if pipe.From.Y == -1 || pipe.To.Y == -1 {
					isInside = !isInside
				}
			} else if isInside {
				enclosed++
			}
		}
	}

	return enclosed
}

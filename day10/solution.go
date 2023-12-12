package day10

import "github.com/kprav33n/aoc23/geometry"

// Pipe represents a pipe.
type Pipe struct {
	From geometry.Point
	To   geometry.Point
	Char string
}

var pipeMapping = map[rune]Pipe{
	'|': {geometry.Point{X: 0, Y: -1}, geometry.Point{X: 0, Y: 1}, "|"},
	'-': {geometry.Point{X: -1, Y: 0}, geometry.Point{X: 1, Y: 0}, "-"},
	'L': {geometry.Point{X: 0, Y: -1}, geometry.Point{X: 1, Y: 0}, "L"},
	'J': {geometry.Point{X: 0, Y: -1}, geometry.Point{X: -1, Y: 0}, "J"},
	'7': {geometry.Point{X: -1, Y: 0}, geometry.Point{X: 0, Y: 1}, "7"},
	'F': {geometry.Point{X: 0, Y: 1}, geometry.Point{X: 1, Y: 0}, "F"},
	'.': {geometry.Point{X: 0, Y: 0}, geometry.Point{X: 0, Y: 0}, "."},
	'S': {geometry.Point{X: 0, Y: 0}, geometry.Point{X: 0, Y: 0}, "S"},
}

// NewPipe returns a new pipe.
func NewPipe(r rune) Pipe {
	return pipeMapping[r]
}

// Sketch represents a sketch.
type Sketch struct {
	Start geometry.Point
	Pipes [][]Pipe
}

// NewSketch returns a new sketch.
func NewSketch(input []string) Sketch {
	sketch := Sketch{}
	var start geometry.Point
	for _, line := range input {
		var row []Pipe
		for _, r := range line {
			if r == 'S' {
				start = geometry.Point{X: len(row), Y: len(sketch.Pipes)}
			}
			row = append(row, NewPipe(r))
		}
		sketch.Pipes = append(sketch.Pipes, row)
	}
	sketch.Start = start
	return sketch
}

// StartPipe returns the pipe at start.
func PipeForDiff(x, y geometry.Point) Pipe {
	for _, pipe := range pipeMapping {
		if (pipe.From == x && pipe.To == y) || (pipe.From == y && pipe.To == x) {
			return pipe
		}
	}
	return pipeMapping['.']
}

// ConnectedToStart returns the pipes that are connected to start.
func (sketch *Sketch) ConnectedToStart() ([]geometry.Point, []geometry.Point) {
	candidates := []geometry.Point{}
	diffs := []geometry.Point{}
	minX, minY := 0, 0
	maxX, maxY := len(sketch.Pipes[0]), len(sketch.Pipes)
	offsets := []geometry.Point{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
	}
	for _, offset := range offsets {
		p := sketch.Start.Add(offset)
		if p.X < minX || p.X >= maxX || p.Y < minY || p.Y >= maxY {
			continue
		}
		diff := sketch.Start.Subtract(p)
		pipe := sketch.Pipes[p.Y][p.X]
		if pipe.From == diff || pipe.To == diff {
			candidates = append(candidates, p)
			diffs = append(diffs, diff)
		}
	}

	startPipe := PipeForDiff(candidates[0].Subtract(sketch.Start), candidates[1].Subtract(sketch.Start))
	sketch.Pipes[sketch.Start.Y][sketch.Start.X] = startPipe

	return candidates, diffs
}

// BackToStart returns the number of steps back to start.
func (sketch *Sketch) BackToStart(p geometry.Point, diff geometry.Point, acc int) int {
	if p == sketch.Start {
		return acc
	}
	pipe := sketch.Pipes[p.Y][p.X]
	var next geometry.Point
	if diff == pipe.From {
		next = p.Add(pipe.To)
	} else {
		next = p.Add(pipe.From)
	}
	return sketch.BackToStart(next, p.Subtract(next), acc+1)
}

// PathToStart returns the path back to start.
func (sketch *Sketch) PathToStart(p geometry.Point, diff geometry.Point, acc []geometry.Point) []geometry.Point {
	if p == sketch.Start {
		return acc
	}
	pipe := sketch.Pipes[p.Y][p.X]
	var next geometry.Point
	if diff == pipe.From {
		next = p.Add(pipe.To)
	} else {
		next = p.Add(pipe.From)
	}
	return sketch.PathToStart(next, p.Subtract(next), append(acc, next))
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
	acc := []geometry.Point{neighbors[0]}
	path := sketch.PathToStart(neighbors[0], diffs[0], acc)
	pathMap := make(map[geometry.Point]struct{})
	for _, p := range path {
		pathMap[p] = struct{}{}
	}

	enclosed := 0
	isInside := false
	for y, row := range sketch.Pipes {
		for x, pipe := range row {
			if _, ok := pathMap[geometry.Point{X: x, Y: y}]; ok {
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

package day16

import (
	"github.com/kprav33n/aoc23/geometry"
)

// NumEnergizedTiles returns the number of energized tiles.
func NumEnergizedTiles(input []string) int {
	c := NewContraption(input)
	return c.NumEnergizedTiles(Beam{
		Position:  geometry.Point{X: 0, Y: 0},
		Direction: Right,
	})
}

// MaxEnergizedTiles returns the maximum number of energized tiles.
func MaxEnergizedTiles(input []string) int {
	c := NewContraption(input)
	return c.MaxEnergizedTiles()
}

// Directions for moving around the contraption.
var (
	Right = geometry.Point{X: 1, Y: 0}
	Left  = geometry.Point{X: -1, Y: 0}
	Up    = geometry.Point{X: 0, Y: -1}
	Down  = geometry.Point{X: 0, Y: 1}
)

// Beam represents a beam of light traveling through the contraption.
type Beam struct {
	Position  geometry.Point
	Direction geometry.Point
}

// Contraption represents a contraption that can be energized.
type Contraption struct {
	NumRows int
	NumCols int
	Grids   [][]rune
}

// NewContraption returns a new contraption.
func NewContraption(input []string) *Contraption {
	c := &Contraption{
		NumRows: len(input),
		NumCols: len(input[0]),
		Grids:   make([][]rune, len(input)),
	}
	for i, row := range input {
		c.Grids[i] = make([]rune, len(row))
		for j, r := range row {
			c.Grids[i][j] = r
		}
	}
	return c
}

// NextBeams returns the next beams for the given beam.
func (c *Contraption) NextBeams(b Beam) []Beam {
	var result []Beam
	var directions []geometry.Point

	switch c.Grids[b.Position.Y][b.Position.X] {
	case '.':
		directions = append(directions, b.Direction)

	case '/':
		reflections := map[geometry.Point]geometry.Point{
			Right: Up,
			Left:  Down,
			Up:    Right,
			Down:  Left,
		}
		directions = append(directions, reflections[b.Direction])

	case '\\':
		reflections := map[geometry.Point]geometry.Point{
			Right: Down,
			Left:  Up,
			Up:    Left,
			Down:  Right,
		}
		directions = append(directions, reflections[b.Direction])

	case '|':
		splits := map[geometry.Point][]geometry.Point{
			Right: {Up, Down},
			Left:  {Up, Down},
			Up:    {Up},
			Down:  {Down},
		}
		directions = splits[b.Direction]

	case '-':
		splits := map[geometry.Point][]geometry.Point{
			Right: {Right},
			Left:  {Left},
			Up:    {Left, Right},
			Down:  {Left, Right},
		}
		directions = splits[b.Direction]
	}

	for _, direction := range directions {
		position := b.Position.Add(direction)
		if position.X < 0 || position.X >= c.NumCols ||
			position.Y < 0 || position.Y >= c.NumRows {
			continue
		}
		result = append(result, Beam{
			Position:  position,
			Direction: direction,
		})
	}

	return result
}

// NumEnergizedTiles returns the number of energized tiles in the contraption.
func (c *Contraption) NumEnergizedTiles(start Beam) int {
	energizedTiles := make(map[geometry.Point]bool)
	seenBeams := make(map[Beam]bool)
	beams := []Beam{start}

	for len(beams) > 0 {
		var nextBeams []Beam
		for _, beam := range beams {
			energizedTiles[beam.Position] = true
			for _, nextBeam := range c.NextBeams(beam) {
				if !seenBeams[nextBeam] {
					nextBeams = append(nextBeams, nextBeam)
					seenBeams[nextBeam] = true
				}
			}
		}
		beams = nextBeams
	}

	return len(energizedTiles)
}

// MaxEnergizedTiles returns the maximum number of energized tiles in the contraption.
func (c *Contraption) MaxEnergizedTiles() int {
	result := 0
	attemptStart := func(start Beam) {
		n := c.NumEnergizedTiles(start)
		if n > result {
			result = n
		}
	}

	for i := 0; i < c.NumRows; i++ {
		attemptStart(Beam{
			Position:  geometry.Point{X: 0, Y: i},
			Direction: Right,
		})
		attemptStart(Beam{
			Position:  geometry.Point{X: c.NumCols - 1, Y: i},
			Direction: Left,
		})
	}

	for i := 0; i < c.NumCols; i++ {
		attemptStart(Beam{
			Position:  geometry.Point{X: i, Y: 0},
			Direction: Down,
		})
		attemptStart(Beam{
			Position:  geometry.Point{X: i, Y: c.NumRows - 1},
			Direction: Up,
		})
	}

	return result
}

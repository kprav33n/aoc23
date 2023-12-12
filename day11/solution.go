package day11

type Point struct {
	X int
	Y int
}

// ManhattanDistance returns the Manhattan distance between two points.
func (p *Point) ManhattanDistance(q Point) int {
	return abs(p.X-q.X) + abs(p.Y-q.Y)
}

// Add returns the sum of two points.
func (p *Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// SumOfShortestPaths returns the sum of the shortest paths.
func SumOfShortestPaths(input []string) int {
	galaxies := []Point{}
	numRows := len(input)
	numCols := len(input[0])
	rowsWithGalaxies := make([]bool, numRows)
	colsWithGalaxies := make([]bool, numCols)
	for y, line := range input {
		for x, r := range line {
			if r == '#' {
				galaxies = append(galaxies, Point{x, y})
				rowsWithGalaxies[y] = true
				colsWithGalaxies[x] = true
			}
		}
	}
	rowOffsets := make([]int, numRows)
	for i := 0; i < numRows-1; i++ {
		if !rowsWithGalaxies[i] {
			rowOffsets[i+1] = rowOffsets[i] + 1
		} else {
			rowOffsets[i+1] = rowOffsets[i]
		}
	}

	colOffsets := make([]int, numCols)
	for i := 0; i < numCols-1; i++ {
		if !colsWithGalaxies[i] {
			colOffsets[i+1] = colOffsets[i] + 1
		} else {
			colOffsets[i+1] = colOffsets[i]
		}
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			p1 := galaxies[i].Add(Point{colOffsets[galaxies[i].X], rowOffsets[galaxies[i].Y]})
			p2 := galaxies[j].Add(Point{colOffsets[galaxies[j].X], rowOffsets[galaxies[j].Y]})
			sum += p1.ManhattanDistance(p2)
		}
	}
	return sum
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

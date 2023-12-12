package geometry

// Point represents a point.
type Point struct {
	X int
	Y int
}

// Add returns the sum of two points.
func (p *Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Subtract returns the difference of two points.
func (p *Point) Subtract(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// ManhattanDistance returns the Manhattan distance between two points.
func (p *Point) ManhattanDistance(q Point) int {
	return abs(p.X-q.X) + abs(p.Y-q.Y)
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

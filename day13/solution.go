package day13

// Pattern represents a pattern of ash and rocks.
type Pattern struct {
	Values [][]rune
}

// NewPattern creates a new pattern from a string.
func NewPattern(input []string) *Pattern {
	p := &Pattern{}
	for _, line := range input {
		p.Values = append(p.Values, []rune(line))
	}
	return p
}

// ReflectionLine returns the row-wise or column-wise reflection line.
func (p *Pattern) ReflectionLine() (int, int) {
	numRows := len(p.Values)
	numCols := len(p.Values[0])

	for i := 1; i < numRows; i++ {
		reflection := true
	outerRow:
		for j := 0; j < i; j++ {
			if i+j >= numRows || i-j-1 < 0 {
				break
			}

			for k := 0; k < numCols; k++ {
				if p.Values[i+j][k] != p.Values[i-j-1][k] {
					reflection = false
					break outerRow
				}
			}
		}
		if reflection {
			return i, 0
		}
	}

	for i := 1; i < numCols; i++ {
		reflection := true
	outerCol:
		for j := 0; j < i; j++ {
			if i+j >= numCols || i-j-1 < 0 {
				break
			}

			for k := 0; k < numRows; k++ {
				if p.Values[k][i+j] != p.Values[k][i-j-1] {
					reflection = false
					break outerCol
				}
			}
		}
		if reflection {
			return 0, i
		}
	}

	return 0, 0
}

// SummaryNumber returns the number after summarizing the notes.
func SummaryNumber(input []string) int {
	markers := []int{-1}
	for i, line := range input {
		if line == "" {
			markers = append(markers, i)
		}
	}
	markers = append(markers, len(input))

	result := 0
	for i := 0; i < len(markers)-1; i++ {
		pattern := NewPattern(input[markers[i]+1 : markers[i+1]])
		row, col := pattern.ReflectionLine()
		result += 100 * row
		result += col
	}
	return result
}

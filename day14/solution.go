package day14

// TotalLoad returns the total load on the north support beams.
func TotalLoad(input []string) int {
	m := NewMap(input)
	m.TiltToNorth()
	return m.TotalLoad()
}

// Map reperesents the map of the rocks and empty spaces.
type Map struct {
	NumRows int
	NumCols int
	Cells   [][]rune
}

// NewMap returns a new map from the input.
func NewMap(input []string) *Map {
	m := &Map{
		NumRows: len(input),
		NumCols: len(input[0]),
		Cells:   make([][]rune, len(input)),
	}
	for i, row := range input {
		m.Cells[i] = make([]rune, len(row))
		for j, ch := range row {
			m.Cells[i][j] = ch
		}
	}
	return m
}

// TiltToNorth tilts the map to the north.
func (m *Map) TiltToNorth() {
	for c := 0; c < m.NumRows; c++ {
		for i := 0; i < m.NumRows-1; i++ {
			for j := 0; j < m.NumCols; j++ {
				if m.Cells[i][j] == '.' && m.Cells[i+1][j] == 'O' {
					m.Cells[i][j] = m.Cells[i+1][j]
					m.Cells[i+1][j] = '.'
				}
			}
		}
	}
}

// TotalLoad returns the total load on the north support beams.
func (m *Map) TotalLoad() int {
	result := 0
	for i := 0; i < m.NumRows; i++ {
		load := m.NumRows - i
		for j := 0; j < m.NumCols; j++ {
			if m.Cells[i][j] == 'O' {
				result += load
			}
		}
	}
	return result
}

// String returns the string representation of the map.
func (m *Map) String() string {
	var s string
	for _, row := range m.Cells {
		s += string(row) + "\n"
	}
	return s
}

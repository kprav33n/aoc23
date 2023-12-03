package day03

import (
	"fmt"
	"strconv"
	"unicode"
)

// NumberLocation represents the location of a number in the schematic.
type NumberLocation struct {
	Row         int
	StartColumn int
	EndColumn   int
}

// NumberLocations returns the location of numbers in the schematic.
func NumberLocations(schematic [][]rune) []NumberLocation {
	var locations []NumberLocation
	startCol, endCol := -1, -1
	for row, line := range schematic {
		if startCol != -1 {
			locations = append(locations, NumberLocation{Row: row - 1, StartColumn: startCol, EndColumn: endCol})
			startCol, endCol = -1, -1
		}
		for column, char := range line {
			if char >= '0' && char <= '9' {
				if startCol == -1 {
					startCol = column
				}
				endCol = column
			} else {
				if startCol != -1 {
					locations = append(locations, NumberLocation{Row: row, StartColumn: startCol, EndColumn: endCol})
					startCol, endCol = -1, -1
				}
			}
		}
	}
	return locations
}

// IsPartNumber returns true if the given number is a part number.
func IsPartNumber(schematic [][]rune, location NumberLocation) bool {
	numRows := len(schematic)
	numCols := len(schematic[0])
	isAdjacentToSymbol := func(row, col int) bool {
		if row < 0 || row >= numRows || col < 0 || col >= numCols {
			return false
		}
		r := schematic[row][col]
		return !unicode.IsNumber(r) && r != '.'
	}
	for col := location.StartColumn; col <= location.EndColumn; col++ {
		if isAdjacentToSymbol(location.Row-1, col-1) ||
			isAdjacentToSymbol(location.Row-1, col) ||
			isAdjacentToSymbol(location.Row-1, col+1) ||
			isAdjacentToSymbol(location.Row, col-1) ||
			isAdjacentToSymbol(location.Row, col+1) ||
			isAdjacentToSymbol(location.Row+1, col-1) ||
			isAdjacentToSymbol(location.Row+1, col) ||
			isAdjacentToSymbol(location.Row+1, col+1) {
			return true
		}
	}
	return false
}

// ExtractNumber returns the number at the given location.
func ExtractNumber(schematic [][]rune, location NumberLocation) int {
	s := string(schematic[location.Row][location.StartColumn : location.EndColumn+1])
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse number at %d, %d-%d: %s", location.Row, location.StartColumn, location.EndColumn, s))
	}
	return int(value)
}

// SumOfPartNumbers returns the sum of all part numbers in the schematic.
func SumOfPartNumbers(schematic [][]rune) int {
	locations := NumberLocations(schematic)
	sum := 0
	for _, location := range locations {
		if IsPartNumber(schematic, location) {
			sum += ExtractNumber(schematic, location)
		}
	}
	return sum
}

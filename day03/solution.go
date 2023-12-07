package day03

import (
	"unicode"

	"github.com/kprav33n/aoc23/strconv"
)

// NumberLocation represents the location of a number in the schematic.
type NumberLocation struct {
	Row         int
	StartColumn int
	EndColumn   int
}

// Location represents a location in the schematic.
type Location struct {
	Row    int
	Column int
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
	return strconv.MustAtoi(s)
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

// GearLocations returns the location of gears in the schematic.
func GearLocations(schematic [][]rune) []Location {
	var locations []Location
	for row, line := range schematic {
		for column, char := range line {
			if char == '*' {
				locations = append(locations, Location{Row: row, Column: column})
			}
		}
	}
	return locations
}

// IsAdjacent returns if the gear location is adjacent to the number location.
func IsAdjacent(schematic [][]rune, gear Location, number NumberLocation) bool {
	numRows := len(schematic)
	numCols := len(schematic[0])
	isAdjacentToGear := func(row, col int) bool {
		if row < 0 || row >= numRows || col < 0 || col >= numCols {
			return false
		}
		return row == gear.Row && col == gear.Column
	}
	for col := number.StartColumn; col <= number.EndColumn; col++ {
		if isAdjacentToGear(number.Row-1, col-1) ||
			isAdjacentToGear(number.Row-1, col) ||
			isAdjacentToGear(number.Row-1, col+1) ||
			isAdjacentToGear(number.Row, col-1) ||
			isAdjacentToGear(number.Row, col+1) ||
			isAdjacentToGear(number.Row+1, col-1) ||
			isAdjacentToGear(number.Row+1, col) ||
			isAdjacentToGear(number.Row+1, col+1) {
			return true
		}
	}
	return false
}

// SumOfGearRatios returns the sum of all gear ratios in the schematic.
func SumOfGearRatios(schematic [][]rune) int {
	gearLocations := GearLocations(schematic)
	numberLocations := NumberLocations(schematic)
	sum := 0
	for _, g := range gearLocations {
		numAdjacent := 0
		gearRatio := 1
		for _, n := range numberLocations {
			if IsAdjacent(schematic, g, n) {
				numAdjacent++
				gearRatio *= ExtractNumber(schematic, n)
			}
		}
		if numAdjacent == 2 {
			sum += gearRatio
		}
	}
	return sum
}

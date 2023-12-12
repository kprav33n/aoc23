package day11

import "github.com/kprav33n/aoc23/geometry"

// SumOfShortestPaths returns the sum of the shortest paths.
func SumOfShortestPaths(input []string, expansionFactor int) int {
	galaxies := []geometry.Point{}
	numRows := len(input)
	numCols := len(input[0])
	rowsWithGalaxies := make([]bool, numRows)
	colsWithGalaxies := make([]bool, numCols)
	for y, line := range input {
		for x, r := range line {
			if r == '#' {
				galaxies = append(galaxies, geometry.Point{X: x, Y: y})
				rowsWithGalaxies[y] = true
				colsWithGalaxies[x] = true
			}
		}
	}

	rowOffsets := make([]int, numRows)
	for i := 0; i < numRows-1; i++ {
		if !rowsWithGalaxies[i] {
			rowOffsets[i+1] = rowOffsets[i] + expansionFactor - 1
		} else {
			rowOffsets[i+1] = rowOffsets[i]
		}
	}

	colOffsets := make([]int, numCols)
	for i := 0; i < numCols-1; i++ {
		if !colsWithGalaxies[i] {
			colOffsets[i+1] = colOffsets[i] + expansionFactor - 1
		} else {
			colOffsets[i+1] = colOffsets[i]
		}
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			p1 := galaxies[i].Add(geometry.Point{
				X: colOffsets[galaxies[i].X],
				Y: rowOffsets[galaxies[i].Y],
			})
			p2 := galaxies[j].Add(geometry.Point{
				X: colOffsets[galaxies[j].X],
				Y: rowOffsets[galaxies[j].Y],
			})
			sum += p1.ManhattanDistance(p2)
		}
	}
	return sum
}

package day03_test

import (
	"testing"

	"infzen.com/aoc23/day03"
)

func TestSumOfPartNumbers(t *testing.T) {
	schematic := [][]rune{
		[]rune("467..114.."),
		[]rune("...*......"),
		[]rune("..35..633."),
		[]rune("......#..."),
		[]rune("617*......"),
		[]rune(".....+.58."),
		[]rune("..592....."),
		[]rune("......755."),
		[]rune("...$.*...."),
		[]rune(".664.598.."),
	}
	expected := 4361
	actual := day03.SumOfPartNumbers(schematic)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

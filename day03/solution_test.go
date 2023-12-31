package day03_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day03"
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

func TestSumOfGearRatios(t *testing.T) {
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
	expected := 467835
	actual := day03.SumOfGearRatios(schematic)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

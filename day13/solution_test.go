package day13_test

import (
	"fmt"
	"testing"

	"github.com/kprav33n/aoc23/day13"
)

func TestReflectionLine(t *testing.T) {
	tests := []struct {
		pattern *day13.Pattern
		row     int
		col     int
	}{
		{
			day13.NewPattern([]string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			}), 0, 5},
		{
			day13.NewPattern([]string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			}), 4, 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case#%d", i), func(t *testing.T) {
			row, col := test.pattern.ReflectionLine()
			if row != test.row || col != test.col {
				t.Errorf("Expected row = %d, col = %d, got row = %d, col = %d", test.row, test.col, row, col)
			}
		})
	}
}

func TestSummaryNumber(t *testing.T) {
	input := []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
		"",
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
	got := day13.SummaryNumber(input)
	want := 405
	if got != want {
		t.Errorf("SummaryNumber() = %d, want %d", got, want)
	}
}

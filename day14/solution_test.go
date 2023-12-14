package day14_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day14"
)

func TestTotalLoad(t *testing.T) {
	input := []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
	want := 136
	got := day14.TotalLoad(input)
	if got != want {
		t.Errorf("TotalLoad(%q) = %d, want %d", input, got, want)
	}
}

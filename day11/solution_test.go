package day11_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day11"
)

func TestSumOfShortestPaths(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	want := 374
	got := day11.SumOfShortestPaths(input)
	if got != want {
		t.Errorf("SumOfShortestPaths(%v) = %v; want %v", input, got, want)
	}
}

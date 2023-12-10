package day10_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day10"
)

func TestStepsToFarthestPoint(t *testing.T) {
	input := []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}
	want := 8
	got := day10.StepsToFarthestPoint(input)
	if got != want {
		t.Errorf("StepsToFarthestPoint(%v) = %v; want %v", input, got, want)
	}
}

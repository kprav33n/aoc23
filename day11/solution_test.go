package day11_test

import (
	"fmt"
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
	tests := []struct {
		expansionFactor int
		want            int
	}{
		{2, 374},
		{10, 1030},
		{100, 8410},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("expansionFactor=%d", test.expansionFactor), func(t *testing.T) {
			got := day11.SumOfShortestPaths(input, test.expansionFactor)
			if got != test.want {
				t.Errorf("SumOfShortestPaths(%v) = %v; want %v", input, got, test.want)
			}
		})
	}
}

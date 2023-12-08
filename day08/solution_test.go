package day08_test

import (
	"fmt"
	"testing"

	"github.com/kprav33n/aoc23/day08"
)

func TestNumStepsToDestination(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			}, 2,
		},
		{
			[]string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			}, 6,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d_steps", test.expected), func(t *testing.T) {
			actual := day08.NumStepsToDestination(test.input)
			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}

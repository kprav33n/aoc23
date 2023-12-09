package day09_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day09"
)

func TestExtrapolateHistory(t *testing.T) {
	tests := []struct {
		history string
		left    int
		right   int
	}{
		{"0 3 6 9 12 15", -3, 18},
		{"1 3 6 10 15 21", 0, 28},
		{"10 13 16 21 30 45", 5, 68},
	}

	for _, test := range tests {
		t.Run(test.history, func(t *testing.T) {
			gotLeft, gotRight := day09.ExtrapolateHistory(test.history)
			if gotLeft != test.left || gotRight != test.right {
				t.Errorf("ExtrapolateHistory(%q) = (%d, %d), want (%d, %d)", test.history, gotLeft, gotRight, test.left, test.right)
			}
		})
	}
}

func TestSumOfExtrapolatedValues(t *testing.T) {
	input := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	wantLeft := 2
	wantRight := 114
	left, right := day09.SumOfExtrapolatedValues(input)
	if left != wantLeft || right != wantRight {
		t.Errorf("SumOfExtrapolatedValues(%q) = (%d, %d), want (%d, %d)", input, left, right, wantLeft, wantRight)
	}
}

package day09_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day09"
)

func TestNextValueInHistory(t *testing.T) {
	tests := []struct {
		history string
		want    int
	}{
		{"0 3 6 9 12 15", 18},
		{"1 3 6 10 15 21", 28},
		{"10 13 16 21 30 45", 68},
		{"-1 9 31 63 102 139 150 87 -129 -609 -1472 -2740 -4028 -3797 2219 24016 82639 219061 508309 1081061 2155578", 0},
	}

	for _, test := range tests {
		t.Run(test.history, func(t *testing.T) {
			got := day09.NextValueInHistory(test.history)
			if got != test.want {
				t.Errorf("NextValueInHistory(%q) = %d, want %d", test.history, got, test.want)
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
	want := 114
	got := day09.SumOfExtrapolatedValues(input)
	if got != want {
		t.Errorf("SumOfExtrapolatedValues(%q) = %d, want %d", input, got, want)
	}
}

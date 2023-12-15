package day15_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day15"
)

func TestSumOfStepHashes(t *testing.T) {
	input := []string{
		"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
	}
	want := 1320
	got := day15.SumOfStepHashes(input)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestHash(t *testing.T) {
	tests := []struct {
		step day15.Step
		want int
	}{
		{step: day15.Step([]rune("rn=1")), want: 30},
		{step: day15.Step([]rune("cm-")), want: 253},
		{step: day15.Step([]rune("qp=3")), want: 97},
		{step: day15.Step([]rune("cm=2")), want: 47},
		{step: day15.Step([]rune("qp-")), want: 14},
		{step: day15.Step([]rune("pc=4")), want: 180},
		{step: day15.Step([]rune("ot=9")), want: 9},
		{step: day15.Step([]rune("ab=5")), want: 197},
		{step: day15.Step([]rune("pc-")), want: 48},
		{step: day15.Step([]rune("pc=6")), want: 214},
		{step: day15.Step([]rune("ot=7")), want: 231},
	}

	for _, tt := range tests {
		t.Run(string(tt.step), func(t *testing.T) {
			got := tt.step.Hash()
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

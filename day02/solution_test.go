package day02_test

import (
	"fmt"
	"testing"

	"infzen.com/aoc23/day02"
)

func TestNewCombination(t *testing.T) {
	tests := []struct {
		input string
		want  day02.Combination
	}{
		{"3 blue, 4 red", day02.Combination{Red: 4, Green: 0, Blue: 3}},
		{"1 red, 2 green", day02.Combination{Red: 1, Green: 2, Blue: 0}},
		{"2 green", day02.Combination{Red: 0, Green: 2, Blue: 0}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := day02.NewCombination(tt.input)
			if got != tt.want {
				t.Errorf("NewCombination(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPossibleCombination(t *testing.T) {
	tests := []struct {
		combination day02.Combination
		constraint  day02.Combination
		want        bool
	}{
		{day02.Combination{Red: 0, Green: 0, Blue: 0}, day02.Combination{Red: 1, Green: 1, Blue: 1}, true},
		{day02.Combination{Red: 1, Green: 1, Blue: 1}, day02.Combination{Red: 1, Green: 1, Blue: 1}, true},
		{day02.Combination{Red: 2, Green: 2, Blue: 2}, day02.Combination{Red: 1, Green: 1, Blue: 1}, false},
		{day02.Combination{Red: 1, Green: 2, Blue: 3}, day02.Combination{Red: 1, Green: 1, Blue: 1}, false},
		{day02.Combination{Red: 1, Green: 2, Blue: 3}, day02.Combination{Red: 2, Green: 2, Blue: 2}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v/%v", tt.combination, tt.constraint), func(t *testing.T) {
			got := day02.IsPossibleCombination(tt.combination, tt.constraint)
			if got != tt.want {
				t.Errorf("IsPossibleCombination(%v, %v) = %t, want %t", tt.combination, tt.constraint, got, tt.want)
			}
		})
	}
}

func TestParseGame(t *testing.T) {
	tests := []struct {
		input            string
		wantID           int
		wantCombinations []day02.Combination
	}{
		{"Game 1: 1 red, 1 green; 2 red, 2 green", 1, []day02.Combination{
			{Red: 1, Green: 1, Blue: 0},
			{Red: 2, Green: 2, Blue: 0},
		}},
		{"Game 2: 1 red, 1 green, 1 blue; 2 red, 2 green, 2 blue", 2, []day02.Combination{
			{Red: 1, Green: 1, Blue: 1},
			{Red: 2, Green: 2, Blue: 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			gotID, gotCombinations := day02.ParseGame(tt.input)
			if gotID != tt.wantID {
				t.Errorf("ParseGame(%q) = %d, want %d", tt.input, gotID, tt.wantID)
			}
			if len(gotCombinations) != len(tt.wantCombinations) {
				t.Errorf("ParseGame(%q) = %d combinations, want %d", tt.input, len(gotCombinations), len(tt.wantCombinations))
			}
			for i := range gotCombinations {
				if gotCombinations[i] != tt.wantCombinations[i] {
					t.Errorf("ParseGame(%q) = %v, want %v", tt.input, gotCombinations[i], tt.wantCombinations[i])
				}
			}
		})
	}
}

func TestSumOfPossibleGameIDs(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	constraint := day02.Combination{Red: 12, Green: 13, Blue: 14}
	want := 8
	got := day02.SumOfPossibleGameIDs(input, constraint)
	if got != want {
		t.Errorf("SumOfPossibleGameIDs(%q, %v) = %d, want %d", input, constraint, got, want)
	}
}

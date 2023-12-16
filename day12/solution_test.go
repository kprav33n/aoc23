package day12_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day12"
)

func TestSumOfArrangements(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
	actual := day12.SumOfArrangements(input)
	expected := 21
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSumOfArrangementsX5(t *testing.T) {
	input := []string{
		"???.### 1,1,3",
		".??..??...?##. 1,1,3",
		"?#?#?#?#?#?#?#? 1,3,1,6",
		"????.#...#... 4,1,1",
		"????.######..#####. 1,6,5",
		"?###???????? 3,2,1",
	}
	actual := day12.SumOfArrangementsX5(input)
	expected := 525152
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

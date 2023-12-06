package day06_test

import (
	"testing"

	"infzen.com/aoc23/day06"
)

func TestProductOfNumberOfWays(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	actual := day06.ProductOfNumberOfWays(input)
	expected := 288
	if actual != expected {
		t.Errorf("ProductOfNumberOfWays(%v) = %v; expected %v", input, actual, expected)
	}
}

func TestNumberOfWays(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	actual := day06.NumberOfWays(input)
	expected := 71503
	if actual != expected {
		t.Errorf("ProductOfNumberOfWays(%v) = %v; expected %v", input, actual, expected)
	}
}

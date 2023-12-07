package day07_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day07"
)

func TestTotalWinnings(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	want := 6440
	got := day07.TotalWinnings(input)
	if got != want {
		t.Errorf("TotalWinnings(%v) = %d, want %d", input, got, want)
	}
}

func TestTotalWinningsWithJoker(t *testing.T) {
	input := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	want := 5905
	got := day07.TotalWinningsWithJoker(input)
	if got != want {
		t.Errorf("TotalWinnings(%v) = %d, want %d", input, got, want)
	}
}

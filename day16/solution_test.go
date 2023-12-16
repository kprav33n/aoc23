package day16_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day16"
)

func TestNumEnergizedTiles(t *testing.T) {
	input := []string{
		`.|...\....`,
		`|.-.\.....`,
		`.....|-...`,
		`........|.`,
		`..........`,
		`.........\`,
		`..../.\\..`,
		`.-.-/..|..`,
		`.|....-|.\`,
		`..//.|....`,
	}
	expected := 46
	got := day16.NumEnergizedTiles(input)
	if got != expected {
		t.Errorf("NumEnergizedTiles(%v) = %d, want %d", input, got, expected)
	}
}

func TestMaxEnergizedTiles(t *testing.T) {
	input := []string{
		`.|...\....`,
		`|.-.\.....`,
		`.....|-...`,
		`........|.`,
		`..........`,
		`.........\`,
		`..../.\\..`,
		`.-.-/..|..`,
		`.|....-|.\`,
		`..//.|....`,
	}
	expected := 51
	got := day16.MaxEnergizedTiles(input)
	if got != expected {
		t.Errorf("MaxEnergizedTiles(%v) = %d, want %d", input, got, expected)
	}
}

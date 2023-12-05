package day05_test

import (
	"fmt"
	"testing"

	"infzen.com/aoc23/day05"
)

func TestDestination(t *testing.T) {
	input := []string{
		"50 98 2",
		"52 50 48",
	}
	m := day05.NewMap(input)
	fmt.Printf("%+v\n", m)
	tests := []struct {
		source int
		want   int
	}{
		{source: 0, want: 0},
		{source: 49, want: 49},
		{source: 98, want: 50},
		{source: 99, want: 51},
		{source: 100, want: 100},
		{source: 50, want: 52},
		{source: 51, want: 53},
		{source: 97, want: 99},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.source), func(t *testing.T) {
			got := m.Destination(test.source)
			if got != test.want {
				t.Errorf("Destination(%d) = %d, want %d", test.source, got, test.want)
			}
		})
	}
}

func TestLowestLocation(t *testing.T) {
	input := []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	want := 35
	got := day05.LowestLocation(input)
	if got != want {
		t.Errorf("LowestLocation(input) = %d, want %d", got, want)
	}
}

package day05

import (
	"math"
	"slices"
	"strings"
	"sync"

	"infzen.com/aoc23/strconv"
)

// MapEntry represents a map entry.
type MapEntry struct {
	SourceStart      int
	DestinationStart int
	Length           int
}

// Map represents a source to destination map.
type Map struct {
	Entries []MapEntry
}

// NewMap creates a new map from the given input.
func NewMap(input []string) *Map {
	entries := []MapEntry{}
	for _, line := range input {
		tokens := strings.Fields(line)
		entry := MapEntry{
			SourceStart:      strconv.MustAtoi(tokens[1]),
			DestinationStart: strconv.MustAtoi(tokens[0]),
			Length:           strconv.MustAtoi(tokens[2]),
		}
		entries = append(entries, entry)
	}
	m := &Map{Entries: entries}
	m.SortEntries()
	return m
}

// SortEntries sorts the map entries by source start.
func (m *Map) SortEntries() {
	slices.SortFunc(m.Entries, func(x, y MapEntry) int {
		return x.SourceStart - y.SourceStart
	})
}

// Destination returns the destination for the given source.
func (m *Map) Destination(source int) int {
	for _, entry := range m.Entries {
		if source >= entry.SourceStart && source < entry.SourceStart+entry.Length {
			return entry.DestinationStart + (source - entry.SourceStart)
		}
	}
	return source
}

// LowestLocation returns the lowest location number for the initial seed numbers.
func LowestLocation(input []string) int {
	separators := []int{}
	for i, line := range input {
		if line == "" {
			separators = append(separators, i)
		}
	}

	maps := []*Map{
		NewMap(input[separators[0]+2 : separators[1]]),
		NewMap(input[separators[1]+2 : separators[2]]),
		NewMap(input[separators[2]+2 : separators[3]]),
		NewMap(input[separators[3]+2 : separators[4]]),
		NewMap(input[separators[4]+2 : separators[5]]),
		NewMap(input[separators[5]+2 : separators[6]]),
		NewMap(input[separators[6]+2:]),
	}

	tokens := strings.Split(input[0], ": ")
	fields := strings.Fields(tokens[1])
	seedNumbers := []int{}
	for _, field := range fields {
		seedNumbers = append(seedNumbers, strconv.MustAtoi(field))
	}

	result := math.MaxInt
	for _, seedNumber := range seedNumbers {
		source, destination := seedNumber, math.MaxInt
		for _, m := range maps {
			destination = m.Destination(source)
			source = destination
		}
		if destination < result {
			result = destination
		}
	}

	return result
}

// LowestLocationForRange returns the lowest location number for the initial seed numbers ranges.
func LowestLocationForRange(input []string) int {
	separators := []int{}
	for i, line := range input {
		if line == "" {
			separators = append(separators, i)
		}
	}

	maps := []*Map{
		NewMap(input[separators[0]+2 : separators[1]]),
		NewMap(input[separators[1]+2 : separators[2]]),
		NewMap(input[separators[2]+2 : separators[3]]),
		NewMap(input[separators[3]+2 : separators[4]]),
		NewMap(input[separators[4]+2 : separators[5]]),
		NewMap(input[separators[5]+2 : separators[6]]),
		NewMap(input[separators[6]+2:]),
	}

	var wg sync.WaitGroup
	lowestLocation := func(start, length int, resultCh chan int) {
		defer wg.Done()
		result := math.MaxInt
		for j := start; j < start+length; j++ {
			source, destination := j, math.MaxInt
			for _, m := range maps {
				destination = m.Destination(source)
				source = destination
			}
			if destination < result {
				result = destination
			}
		}
		resultCh <- result
	}

	tokens := strings.Split(input[0], ": ")
	fields := strings.Fields(tokens[1])
	resultCh := make(chan int, len(fields)/2)
	wg.Add(len(fields) / 2)
	for i := 0; i < len(fields); i += 2 {
		start := strconv.MustAtoi(fields[i])
		length := strconv.MustAtoi(fields[i+1])
		go lowestLocation(start, length, resultCh)
	}
	wg.Wait()
	close(resultCh)

	result := math.MaxInt
	for r := range resultCh {
		if r < result {
			result = r
		}
	}
	return result
}

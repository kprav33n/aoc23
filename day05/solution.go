package day05

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type MapEntry struct {
	SourceStart      int
	DestinationStart int
	Length           int
}

type Map struct {
	Entries []MapEntry
}

func NewMap(input []string) *Map {
	entries := []MapEntry{}
	for _, line := range input {
		tokens := strings.Fields(line)
		entry := MapEntry{
			SourceStart:      toInt(tokens[1]),
			DestinationStart: toInt(tokens[0]),
			Length:           toInt(tokens[2]),
		}
		entries = append(entries, entry)
	}
	slices.SortFunc(entries, func(x, y MapEntry) int {
		return x.SourceStart - y.SourceStart
	})
	return &Map{Entries: entries}
}

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
		seedNumbers = append(seedNumbers, toInt(field))
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

	tokens := strings.Split(input[0], ": ")
	fields := strings.Fields(tokens[1])
	result := math.MaxInt
	for i := 0; i < len(fields); i += 2 {
		start := toInt(fields[i])
		length := toInt(fields[i+1])
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
	}

	return result
}

// toInt converts a string to an integer.
func toInt(input string) int {
	n, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("could not convert %s to int", input))
	}
	return n
}

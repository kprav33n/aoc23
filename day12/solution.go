package day12

import (
	"slices"
	"strings"

	istrconv "github.com/kprav33n/aoc23/strconv"
)

// SumOfArrangements returns the sum of the number of arrangements of the given
// input.
func SumOfArrangements(input []string) int {
	sum := 0
	for _, line := range input {
		sum += NumArrangements(line)
	}
	return sum
}

// NumArrangements returns the number of arrangements of the given input.
func NumArrangements(input string) int {
	tokens := strings.Split(input, " ")
	criteria := make([]int, 0)
	for _, field := range strings.Split(tokens[1], ",") {
		criteria = append(criteria, istrconv.MustAtoi(field))
	}

	var bitmask, record uint64
	runes := []rune(tokens[0])
	slices.Reverse(runes)
	length := len(runes)
	for i, r := range runes {
		if r == '#' || r == '.' {
			bitmask |= 1 << i
		}
		if r == '.' {
			record |= 1 << i
		}
	}

	result := 0
	for i := uint64(0); i < 1<<uint64(length); i++ {
		if i&bitmask != record {
			continue
		}

		if slices.Equal(PatternOfZeros(i, length), criteria) {
			result++
		}
	}

	return result
}

// PatternOfZeros returns the pattern of zeros of the given input.
func PatternOfZeros(v uint64, l int) []int {
	result := make([]int, 0)
	numZeros := 0
	for i := l - 1; i >= 0; i-- {
		if v&(1<<uint64(i)) == 0 {
			numZeros++
		} else {
			if numZeros > 0 {
				result = append(result, numZeros)
				numZeros = 0
			}
		}
	}
	if numZeros > 0 {
		result = append(result, numZeros)
	}
	return result
}

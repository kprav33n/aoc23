package day12

import (
	"strings"

	istrconv "github.com/kprav33n/aoc23/strconv"
)

// SumOfArrangements returns the sum of the number of arrangements of the given
// input.
func SumOfArrangements(input []string) int {
	sum := 0
	for _, line := range input {
		sum += NumArrangementsScaled(line, 1)
	}
	return sum
}

// SumOfArrangementsX5 returns the sum of the number of arrangements of the
// given input scaled by 5.
func SumOfArrangementsX5(input []string) int {
	sum := 0
	for _, line := range input {
		sum += NumArrangementsScaled(line, 5)
	}
	return sum
}

// NumArrangementsScaled returns the number of arrangements of the given input
// scaled by the given factor.
func NumArrangementsScaled(input string, factor int) int {
	tokens := strings.Split(input, " ")
	repeater := make([]string, factor)
	for i := range repeater {
		repeater[i] = tokens[0]
	}
	condition := []rune(strings.Join(repeater, "?"))

	for i := range repeater {
		repeater[i] = tokens[1]
	}
	criteria := make([]int, 0)
	for _, field := range strings.Split(strings.Join(repeater, ","), ",") {
		criteria = append(criteria, istrconv.MustAtoi(field))
	}

	return newCache().numArrangements(condition, criteria, 0, 0, 0)
}

// cache represents a cache of triplets.
type cache struct {
	data map[triplet]int
}

// newCache returns a new cache.
func newCache() *cache {
	return &cache{make(map[triplet]int)}
}

// get returns the value for the given triplet.
func (c *cache) get(x, y, z int) (int, bool) {
	result, ok := c.data[triplet{x, y, z}]
	return result, ok
}

// set sets the value for the given triplet.
func (c *cache) set(x, y, z, value int) {
	c.data[triplet{x, y, z}] = value
}

// numArrangements returns the number of arrangements of the given condition
// that satisfy the given criteria.
// Idea courtesy: https://www.youtube.com/watch?v=xTGkP2GNmbQ
func (c *cache) numArrangements(condition []rune, criteria []int,
	conditionOffset, criteriaOffset, groupLen int) int {
	if value, ok := c.get(conditionOffset, criteriaOffset, groupLen); ok {
		return value
	}

	if conditionOffset == len(condition) {
		switch {
		case criteriaOffset == len(criteria) && groupLen == 0:
			return 1
		case criteriaOffset == len(criteria)-1 && criteria[criteriaOffset] == groupLen:
			return 1
		default:
			return 0
		}
	}

	result := 0
	for _, ch := range []rune{'.', '#'} {
		if condition[conditionOffset] != '?' && condition[conditionOffset] != ch {
			continue
		}

		switch {
		case ch == '.' && groupLen == 0:
			result += c.numArrangements(condition, criteria, conditionOffset+1, criteriaOffset, 0)
		case ch == '.' && groupLen > 0 && criteriaOffset < len(criteria) && criteria[criteriaOffset] == groupLen:
			result += c.numArrangements(condition, criteria, conditionOffset+1, criteriaOffset+1, 0)
		case ch == '#':
			result += c.numArrangements(condition, criteria, conditionOffset+1, criteriaOffset, groupLen+1)
		}
	}

	c.set(conditionOffset, criteriaOffset, groupLen, result)
	return result
}

// triplet represents a triplet of integers.
type triplet struct {
	x, y, z int
}

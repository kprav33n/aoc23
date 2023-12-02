package day02

import (
	"fmt"
	"strconv"
	"strings"
)

// Combination represents a combination of red, green and blue cubes.
type Combination struct {
	Red   int
	Blue  int
	Green int
}

// NewCombination parses a string like "3 blue, 4 red" and returns a Combination.
func NewCombination(input string) Combination {
	result := Combination{}
	token := strings.Split(input, ", ")
	for _, t := range token {
		parts := strings.Split(t, " ")
		value, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Sprintf("invalid input: %q, failed to parse %s", input, parts[0]))
		}
		switch parts[1] {
		case "red":
			result.Red = value
		case "green":
			result.Green = value
		case "blue":
			result.Blue = value
		}
	}
	return result
}

// IsPossibleCombination returns true if the combination is allowed.
func IsPossibleCombination(combination Combination, constraint Combination) bool {
	return combination.Red <= constraint.Red &&
		combination.Green <= constraint.Green &&
		combination.Blue <= constraint.Blue
}

// ParseGame parses the game input and returns game ID and slice of combinations.
func ParseGame(input string) (int, []Combination) {
	tokens := strings.Split(input, ": ")

	gameID, err := strconv.Atoi(tokens[0][5:])
	if err != nil {
		panic(fmt.Sprintf("invalid input: %q, failed to parse %s", input, tokens[0][5:]))
	}
	var combinations []Combination
	for _, c := range strings.Split(tokens[1], "; ") {
		combinations = append(combinations, NewCombination(c))
	}
	return gameID, combinations
}

// SumOfPossibleGameIDs returns the sum of all possible game IDs.
func SumOfPossibleGameIDs(input []string, constraint Combination) int {
	var sum int
	for _, line := range input {
		gameID, combinations := ParseGame(line)
		possible := true
		for _, c := range combinations {
			if !IsPossibleCombination(c, constraint) {
				possible = false
				break
			}
		}
		if possible {
			sum += gameID
		}
	}
	return sum
}

package day15

import (
	"fmt"
	"strings"
)

// SumOfStepHashes returns the sum of the hashes of the steps.
func SumOfStepHashes(input []string) int {
	if len(input) != 1 {
		panic(fmt.Sprintf("expected 1 input, got %d", len(input)))
	}

	result := 0
	for _, s := range strings.Split(input[0], ",") {
		result += Step([]rune(s)).Hash()
	}
	return result
}

// Step represents a step in the input.
type Step []rune

// Hash returns a hash of the step.
func (s Step) Hash() int {
	result := 0
	for _, r := range s {
		result += int(r)
		result *= 17
		result %= 256
	}

	return result
}

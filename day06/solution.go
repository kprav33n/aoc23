package day06

import (
	"strings"

	"infzen.com/aoc23/strconv"
)

// distanceTravelled returns the distance travelled in the given time when the
// button is pressed for the given time.
func distanceTravelled(time, button int) int {
	return (time - button) * button
}

// ProductOfNumberOfWays returns the product of the number of ways to win the races.
func ProductOfNumberOfWays(input []string) int {
	tokens := strings.Split(input[0], ":")
	fields := strings.Fields(tokens[1])
	times := make([]int, len(fields))
	for i, field := range fields {
		times[i] = strconv.MustAtoi(field)
	}

	tokens = strings.Split(input[1], ":")
	fields = strings.Fields(tokens[1])
	distances := make([]int, len(fields))
	for i, field := range fields {
		distances[i] = strconv.MustAtoi(field)
	}

	result := 1
	for i := 0; i < len(times); i++ {
		numWays := 0
		for j := 0; j <= times[i]; j++ {
			if distanceTravelled(times[i], j) > distances[i] {
				numWays++
			}
		}
		result *= numWays
	}

	return result
}

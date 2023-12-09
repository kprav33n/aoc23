package day09

import (
	"strings"

	"github.com/kprav33n/aoc23/strconv"
)

// ExtrapolateHistory returns the next value in the history.
func ExtrapolateHistory(history string) (int, int) {
	fields := strings.Fields(history)
	histories := make([][]int, 0)
	histories = append(histories, make([]int, len(fields)))
	j := 0
	for i, field := range fields {
		histories[j][i] = strconv.MustAtoi(field)
	}
	for {
		histories = append(histories, make([]int, len(histories[j])-1))
		k := j + 1
		for i := 0; i < len(histories[j])-1; i++ {
			diff := histories[j][i+1] - histories[j][i]
			histories[k][i] = diff
		}

		allZeros := true
		for _, v := range histories[k][:len(histories[k])-1] {
			if v != 0 {
				allZeros = false
				break
			}
		}
		if allZeros {
			break
		}
		j = k
	}

	left, right := 0, 0
	for j >= 0 {
		left = histories[j][0] - left
		right = histories[j][len(histories[j])-1] + right
		j--
	}

	return left, right
}

// SumOfExtrapolatedValues returns the sum of extrapolated values.
func SumOfExtrapolatedValues(input []string) (int, int) {
	sumLeft, sumRight := 0, 0
	for _, history := range input {
		left, right := ExtrapolateHistory(history)
		sumLeft += left
		sumRight += right
	}
	return sumLeft, sumRight
}

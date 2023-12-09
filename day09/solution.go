package day09

import (
	"fmt"
	"strings"

	"github.com/kprav33n/aoc23/strconv"
)

// NextValueInHistory returns the next value in the history.
func NextValueInHistory(history string) int {
	fields := strings.Fields(history)
	histories := make([][]int, 0)
	histories = append(histories, make([]int, len(fields)+1))
	j := 0
	for i, field := range fields {
		histories[j][i] = strconv.MustAtoi(field)
	}
	for {
		histories = append(histories, make([]int, len(histories[j])-1))
		k := j + 1
		for i := 0; i < len(histories[j])-2; i++ {
			diff := histories[j][i+1] - histories[j][i]
			if diff < 0 {
				diff = -diff
			}
			histories[k][i] = diff
		}

		allZeros := true
		for _, v := range histories[k][:len(histories[k])-1] {
			if v != 0 {
				allZeros = false
				break
			}
		}
		j = k
		if allZeros {
			break
		}
	}
	histories[j][len(histories[j])-1] = 0
	for j > 0 {
		j--
		x := len(histories[j]) - 1
		y := len(histories[j+1]) - 1
		histories[j][x] = histories[j+1][y] + histories[j][x-1]
	}
	for _, history := range histories {
		fmt.Println(history)
	}
	return histories[0][len(histories[0])-1]
}

// SumOfExtrapolatedValues returns the sum of extrapolated values.
func SumOfExtrapolatedValues(input []string) int {
	sum := 0
	for _, history := range input {
		sum += NextValueInHistory(history)
	}
	return sum
}

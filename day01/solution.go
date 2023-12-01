package day01

import (
	"math"
	"strings"
)

// CalibrationValue returns the calibration value for the given input.
func CalibrationValue(input string) int {
	runes := []rune(input)
	first := strings.IndexAny(input, "0123456789")
	last := strings.LastIndexAny(input, "0123456789")
	return int((runes[first]-'0')*10 + runes[last] - '0')
}

// SumOfCalibrationValues returns the sum of calibration values for the given inputs.
func SumOfCalibrationValues(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += CalibrationValue(input)
	}
	return sum
}

var tokens = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// AdjustedCalibrationValue returns the adjusted calibration value for the given input.
func AdjustedCalibrationValue(input string) int {
	runes := []rune(input)
	first := strings.IndexAny(input, "0123456789")
	last := strings.LastIndexAny(input, "0123456789")

	firstTokenIndex, lastTokenIndex := math.MaxInt, math.MinInt
	firstValue, lastValue := 0, 0
	for value, token := range tokens {
		i := strings.Index(input, token)
		if i != -1 && i < firstTokenIndex {
			firstTokenIndex = i
			firstValue = value
		}

		j := strings.LastIndex(input, token)
		if j != -1 && j > lastTokenIndex {
			lastTokenIndex = j
			lastValue = value
		}
	}

	if first != -1 && firstTokenIndex > first {
		firstValue = int(runes[first] - '0')
	}

	if last != -1 && lastTokenIndex < last {
		lastValue = int(runes[last] - '0')
	}

	return firstValue*10 + lastValue
}

// SumOfAdjustedCalibrationValues returns the sum of adjusted calibration values for the given inputs.
func SumOfAdjustedCalibrationValues(inputs []string) int {
	sum := 0
	for _, input := range inputs {
		sum += AdjustedCalibrationValue(input)
	}
	return sum
}

package day01

import (
	"math"
	"strings"
)

// CalibrationValue returns the calibration value for the given input.
func CalibrationValue(input string) int {
	runes := []rune(input)
	var first, last int
	for i := 0; i < len(runes); i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			first = i
			break
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] >= '0' && runes[i] <= '9' {
			last = i
			break
		}
	}

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

var replacementMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// AdjustedCalibrationValue returns the adjusted calibration value for the given input.
func AdjustedCalibrationValue(input string) int {
	runes := []rune(input)
	first := math.MaxInt
	for i := 0; i < len(runes); i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			first = i
			break
		}
	}

	firstTokenIndex := math.MaxInt
	firstToken := ""
	for token := range replacementMap {
		i := strings.Index(input, token)
		if i != -1 && i < firstTokenIndex {
			firstTokenIndex = i
			firstToken = token
		}
	}

	firstValue := 0
	if firstTokenIndex < first {
		firstValue = replacementMap[firstToken]
	} else {
		firstValue = int(runes[first] - '0')
	}

	last := math.MinInt
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] >= '0' && runes[i] <= '9' {
			last = i
			break
		}
	}

	lastTokenIndex := math.MinInt
	lastToken := ""
	for token := range replacementMap {
		i := strings.LastIndex(input, token)
		if i != -1 && i > lastTokenIndex {
			lastTokenIndex = i
			lastToken = token
		}
	}

	lastValue := 0
	if lastTokenIndex > last {
		lastValue = replacementMap[lastToken]
	} else {
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

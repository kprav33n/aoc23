package day01

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

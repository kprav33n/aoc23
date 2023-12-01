package day01_test

import (
	"testing"

	"infzen.com/aoc23/day01"
)

func TestCalibrationValue(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := day01.CalibrationValue(tt.input)
			if got != tt.want {
				t.Errorf("CalibrationValue(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestSumOfCalibrationValues(t *testing.T) {
	tests := []struct {
		inputs []string
		want   int
	}{
		{[]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}, 142},
	}

	for _, tt := range tests {
		t.Run("example", func(t *testing.T) {
			got := day01.SumOfCalibrationValues(tt.inputs)
			if got != tt.want {
				t.Errorf("SumOfCalibrationValues(%q) = %d, want %d", tt.inputs, got, tt.want)
			}
		})
	}
}

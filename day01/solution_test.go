package day01_test

import (
	"testing"

	"github.com/kprav33n/aoc23/day01"
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

func TestAdjustedCalibrationValue(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := day01.AdjustedCalibrationValue(tt.input)
			if got != tt.want {
				t.Errorf("CalibrationValue(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestSumOfAdjustedCalibrationValues(t *testing.T) {
	tests := []struct {
		inputs []string
		want   int
	}{
		{[]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}, 281},
	}

	for _, tt := range tests {
		t.Run("example", func(t *testing.T) {
			got := day01.SumOfAdjustedCalibrationValues(tt.inputs)
			if got != tt.want {
				t.Errorf("SumOfCalibrationValues(%q) = %d, want %d", tt.inputs, got, tt.want)
			}
		})
	}
}

package strconv_test

import (
	"testing"

	"github.com/kprav33n/aoc23/strconv"
)

func TestMustAtoi(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"0", 0},
		{"-1", -1},
		{"1", 1},
		{"-100", -100},
		{"100", 100},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := strconv.MustAtoi(test.input)
			if got != test.want {
				t.Errorf("MustAtoi(%q) = %d, want %d", test.input, got, test.want)
			}
		})
	}
}

func TestMustAtoiPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustAtoi did not panic")
		}
	}()
	strconv.MustAtoi("abc")
}

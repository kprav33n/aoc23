package geometry_test

import (
	"fmt"
	"testing"

	"github.com/kprav33n/aoc23/geometry"
)

func TestPointAdd(t *testing.T) {
	tests := []struct {
		p, q, want geometry.Point
	}{
		{geometry.Point{1, 2}, geometry.Point{3, 4}, geometry.Point{4, 6}},
		{geometry.Point{1, 2}, geometry.Point{-3, -4}, geometry.Point{-2, -2}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v.Add(%v)", test.p, test.q), func(t *testing.T) {
			got := test.p.Add(test.q)
			if got != test.want {
				t.Errorf("%v.Add(%v) = %v; want %v", test.p, test.q, got, test.want)
			}
		})
	}
}

func TestPointSubtract(t *testing.T) {
	tests := []struct {
		p, q, want geometry.Point
	}{
		{geometry.Point{1, 2}, geometry.Point{3, 4}, geometry.Point{-2, -2}},
		{geometry.Point{1, 2}, geometry.Point{-3, -4}, geometry.Point{4, 6}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v.Subtract(%v)", test.p, test.q), func(t *testing.T) {
			got := test.p.Subtract(test.q)
			if got != test.want {
				t.Errorf("%v.Subtract(%v) = %v; want %v", test.p, test.q, got, test.want)
			}
		})
	}
}

func TestPointManhattanDistance(t *testing.T) {
	tests := []struct {
		p, q geometry.Point
		want int
	}{
		{geometry.Point{1, 2}, geometry.Point{3, 4}, 4},
		{geometry.Point{1, 2}, geometry.Point{-3, -4}, 10},
		{geometry.Point{1, 2}, geometry.Point{1, 2}, 0},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v.ManhattanDistance(%v)", test.p, test.q), func(t *testing.T) {
			got := test.p.ManhattanDistance(test.q)
			if got != test.want {
				t.Errorf("%v.ManhattanDistance(%v) = %v; want %v", test.p, test.q, got, test.want)
			}
		})
	}
}

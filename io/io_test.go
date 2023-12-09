package io_test

import (
	_ "embed"
	"slices"
	"strings"
	"testing"

	"github.com/kprav33n/aoc23/io"
)

//go:embed io_test.go
var testFileContent string

func TestReadLines(t *testing.T) {
	got := io.ReadLines("io_test.go")
	want := strings.Split(testFileContent, "\n")
	if !slices.Equal(got, want[:len(want)-1]) {
		t.Errorf("ReadLines(%q) = %q, want %q", "io_test.go", got, want)
	}
}

func TestReadLinesPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ReadLines did not panic")
		}
	}()
	io.ReadLines("abc")
}

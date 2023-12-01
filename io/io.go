package io

import (
	"fmt"
	"os"
)

// ReadLines reads a whole file and returns the content as a slice of strings.
func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	for {
		var line string
		_, err := fmt.Fscanf(file, "%s\n", &line)
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	return lines
}

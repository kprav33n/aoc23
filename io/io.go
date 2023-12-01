package io

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLines reads a whole file and returns the content as a slice of strings.
func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Errorf("error opening file: %v", err))
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("error reading file: %v", err))
	}

	return lines
}

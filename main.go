package main

import (
	"fmt"
	"os"

	"infzen.com/aoc23/day01"
	"infzen.com/aoc23/io"
)

func main() {
	switch os.Args[1] {
	case "day01a":
		input := io.ReadLines("inputs/day01.txt")
		fmt.Println(day01.SumOfCalibrationValues(input))

	case "day01b":
		input := io.ReadLines("inputs/day01.txt")
		fmt.Println(day01.SumOfAdjustedCalibrationValues(input))
	}
}

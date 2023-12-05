package main

import (
	"fmt"
	"os"

	"infzen.com/aoc23/day01"
	"infzen.com/aoc23/day02"
	"infzen.com/aoc23/day03"
	"infzen.com/aoc23/day04"
	"infzen.com/aoc23/day05"
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

	case "day02a":
		input := io.ReadLines("inputs/day02.txt")
		fmt.Println(day02.SumOfPossibleGameIDs(input,
			day02.Combination{Red: 12, Green: 13, Blue: 14}))

	case "day02b":
		input := io.ReadLines("inputs/day02.txt")
		fmt.Println(day02.SumOfPowers(input))

	case "day03a":
		input := io.ReadLines("inputs/day03.txt")
		schematic := [][]rune{}
		for _, line := range input {
			schematic = append(schematic, []rune(line))
		}
		fmt.Println(day03.SumOfPartNumbers(schematic))

	case "day03b":
		input := io.ReadLines("inputs/day03.txt")
		schematic := [][]rune{}
		for _, line := range input {
			schematic = append(schematic, []rune(line))
		}
		fmt.Println(day03.SumOfGearRatios(schematic))

	case "day04a":
		input := io.ReadLines("inputs/day04.txt")
		fmt.Println(day04.TotalPoints(input))

	case "day04b":
		input := io.ReadLines("inputs/day04.txt")
		fmt.Println(day04.TotalCards(input))

	case "day05a":
		input := io.ReadLines("inputs/day05.txt")
		fmt.Println(day05.LowestLocation(input))
	}
}

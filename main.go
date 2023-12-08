package main

import (
	"fmt"
	"os"

	"github.com/kprav33n/aoc23/day01"
	"github.com/kprav33n/aoc23/day02"
	"github.com/kprav33n/aoc23/day03"
	"github.com/kprav33n/aoc23/day04"
	"github.com/kprav33n/aoc23/day05"
	"github.com/kprav33n/aoc23/day06"
	"github.com/kprav33n/aoc23/day07"
	"github.com/kprav33n/aoc23/day08"
	"github.com/kprav33n/aoc23/io"
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

	case "day05b":
		input := io.ReadLines("inputs/day05.txt")
		fmt.Println(day05.LowestLocationForRange(input))

	case "day06a":
		input := io.ReadLines("inputs/day06.txt")
		fmt.Println(day06.ProductOfNumberOfWays(input))

	case "day06b":
		input := io.ReadLines("inputs/day06.txt")
		fmt.Println(day06.NumberOfWays(input))

	case "day07a":
		input := io.ReadLines("inputs/day07.txt")
		fmt.Println(day07.TotalWinnings(input))

	case "day07b":
		input := io.ReadLines("inputs/day07.txt")
		fmt.Println(day07.TotalWinningsWithJoker(input))

	case "day08a":
		input := io.ReadLines("inputs/day08.txt")
		fmt.Println(day08.NumStepsToDestination(input))

	case "day08b":
		input := io.ReadLines("inputs/day08.txt")
		fmt.Println(day08.NumStepsToDestinationForGhost(input))
	}
}

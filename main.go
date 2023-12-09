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
	"github.com/kprav33n/aoc23/day09"
	"github.com/kprav33n/aoc23/io"
)

var solutions = map[string]func(){
	"day01a": func() {
		input := io.ReadLines("inputs/day01.txt")
		fmt.Println(day01.SumOfCalibrationValues(input))
	},

	"day01b": func() {
		input := io.ReadLines("inputs/day01.txt")
		fmt.Println(day01.SumOfAdjustedCalibrationValues(input))
	},
	"day02a": func() {
		input := io.ReadLines("inputs/day02.txt")
		fmt.Println(day02.SumOfPossibleGameIDs(input,
			day02.Combination{Red: 12, Green: 13, Blue: 14}))
	},
	"day02b": func() {
		input := io.ReadLines("inputs/day02.txt")
		fmt.Println(day02.SumOfPowers(input))
	},
	"day03a": func() {
		input := io.ReadLines("inputs/day03.txt")
		schematic := [][]rune{}
		for _, line := range input {
			schematic = append(schematic, []rune(line))
		}
		fmt.Println(day03.SumOfPartNumbers(schematic))
	},
	"day03b": func() {
		input := io.ReadLines("inputs/day03.txt")
		schematic := [][]rune{}
		for _, line := range input {
			schematic = append(schematic, []rune(line))
		}
		fmt.Println(day03.SumOfGearRatios(schematic))
	},
	"day04a": func() {
		input := io.ReadLines("inputs/day04.txt")
		fmt.Println(day04.TotalPoints(input))
	},
	"day04b": func() {
		input := io.ReadLines("inputs/day04.txt")
		fmt.Println(day04.TotalCards(input))
	},
	"day05a": func() {
		input := io.ReadLines("inputs/day05.txt")
		fmt.Println(day05.LowestLocation(input))
	},
	"day05b": func() {
		input := io.ReadLines("inputs/day05.txt")
		fmt.Println(day05.LowestLocationForRange(input))
	},
	"day06a": func() {
		input := io.ReadLines("inputs/day06.txt")
		fmt.Println(day06.ProductOfNumberOfWays(input))
	},
	"day06b": func() {
		input := io.ReadLines("inputs/day06.txt")
		fmt.Println(day06.NumberOfWays(input))
	},
	"day07a": func() {
		input := io.ReadLines("inputs/day07.txt")
		fmt.Println(day07.TotalWinnings(input))
	},
	"day07b": func() {
		input := io.ReadLines("inputs/day07.txt")
		fmt.Println(day07.TotalWinningsWithJoker(input))
	},
	"day08a": func() {
		input := io.ReadLines("inputs/day08.txt")
		fmt.Println(day08.NumStepsToDestination(input))
	},
	"day08b": func() {
		input := io.ReadLines("inputs/day08.txt")
		fmt.Println(day08.NumStepsToDestinationForGhost(input))
	},
	"day09a": func() {
		input := io.ReadLines("inputs/day09.txt")
		_, result := day09.SumOfExtrapolatedValues(input)
		fmt.Println(result)
	},
	"day09b": func() {
		input := io.ReadLines("inputs/day09.txt")
		result, _ := day09.SumOfExtrapolatedValues(input)
		fmt.Println(result)
	},
}

func main() {
	solution, ok := solutions[os.Args[1]]
	if !ok {
		fmt.Println("No solution found for", os.Args[1])
		os.Exit(1)
	}
	solution()
}

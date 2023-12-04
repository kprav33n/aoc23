package day04

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	WinningNumbers map[int]struct{}
	Numbers        []int
}

func NewCard(input string) *Card {
	card := &Card{
		WinningNumbers: make(map[int]struct{}),
		Numbers:        []int{},
	}
	parts := strings.Split(input, ": ")
	parts = strings.Split(parts[1], " | ")
	winningNumbers := strings.Fields(parts[0])
	numbers := strings.Fields(parts[1])
	for _, number := range winningNumbers {
		card.WinningNumbers[toInt(number)] = struct{}{}
	}
	for _, number := range numbers {
		card.Numbers = append(card.Numbers, toInt(number))
	}
	return card
}

func (c *Card) Points() int {
	points := 0
	for _, number := range c.Numbers {
		if _, ok := c.WinningNumbers[number]; ok {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

// TotalPoints returns the total points for the given list of cards.
func TotalPoints(input []string) int {
	points := 0
	for _, line := range input {
		card := NewCard(line)
		points += card.Points()
	}
	return points
}

// toInt converts a string to an integer.
func toInt(input string) int {
	n, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		panic(fmt.Sprintf("could not convert %s to int", input))
	}
	return n
}

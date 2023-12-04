package day04

import (
	"fmt"
	"strconv"
	"strings"
)

// Card represents a scratchcard.
type Card struct {
	WinningNumbers map[int]struct{}
	Numbers        []int
}

// NewCard creates a new card from the given input.
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

// Points returns the winning points for the given card.
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

// NumWins returns the number of wins for the given card.
func (c *Card) NumWins() int {
	numWins := 0
	for _, number := range c.Numbers {
		if _, ok := c.WinningNumbers[number]; ok {
			numWins++
		}
	}
	return numWins
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

// TotalCards returns the total number of cards resulted.
func TotalCards(input []string) int {
	cards := []*Card{}
	for _, line := range input {
		cards = append(cards, NewCard(line))
	}
	numInstances := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		numInstances[i] = 1
	}
	for i := 0; i < len(cards); i++ {
		numWins := cards[i].NumWins()
		for j := 1; j <= numWins; j++ {
			numInstances[i+j] += numInstances[i]
		}
	}
	totalCards := 0
	for i := 0; i < len(cards); i++ {
		totalCards += numInstances[i]
	}
	return totalCards
}

// toInt converts a string to an integer.
func toInt(input string) int {
	n, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("could not convert %s to int", input))
	}
	return n
}

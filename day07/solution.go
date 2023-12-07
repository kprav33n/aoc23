package day07

import (
	"slices"
	"strings"

	"infzen.com/aoc23/strconv"
)

var cardStrength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

// Hand represents a hand of cards.
type Hand struct {
	Cards     []rune
	CardCount map[rune]int
	Bet       int
}

// NewHand returns a new hand from the given cards.
func NewHand(cards []rune, bet int) *Hand {
	cardCount := map[rune]int{}
	for _, card := range cards {
		cardCount[card]++
	}
	return &Hand{
		Cards:     cards,
		CardCount: cardCount,
		Bet:       bet,
	}
}

// Strength returns the strength of the hand.
func (h *Hand) Strength() int {
	counts := []int{}
	for _, count := range h.CardCount {
		counts = append(counts, count)
	}
	slices.Sort(counts)
	slices.Reverse(counts)
	if counts[0] == 5 {
		return 7
	}
	if counts[0] == 4 {
		return 6
	}
	if counts[0] == 3 && counts[1] == 2 {
		return 5
	}
	if counts[0] == 3 {
		return 4
	}
	if counts[0] == 2 && counts[1] == 2 {
		return 3
	}
	if counts[0] == 2 {
		return 2
	}
	return 1
}

// Compare returns the result of comparing the given hands.
func Compare(a, b *Hand) int {
	if a.Strength() > b.Strength() {
		return 1
	}
	if a.Strength() < b.Strength() {
		return -1
	}
	for i := 0; i < len(a.Cards); i++ {
		x := cardStrength[a.Cards[i]]
		y := cardStrength[b.Cards[i]]
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

// TotalWinnings returns the total winnings for the given input.
func TotalWinnings(input []string) int {
	hands := []*Hand{}
	for _, line := range input {
		tokens := strings.Split(line, " ")
		hands = append(hands, NewHand([]rune(tokens[0]), strconv.MustAtoi(tokens[1])))
	}
	slices.SortFunc(hands, func(a, b *Hand) int {
		return Compare(a, b)
	})
	winnings := 0
	for i, hand := range hands {
		winnings += hand.Bet * (i + 1)
	}
	return winnings
}

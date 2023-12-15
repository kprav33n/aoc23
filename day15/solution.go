package day15

import (
	"regexp"
	"strings"

	"github.com/kprav33n/aoc23/strconv"
)

// SumOfStepHashes returns the sum of the hashes of the steps.
func SumOfStepHashes(input []string) int {
	result := 0
	for _, s := range strings.Split(input[0], ",") {
		result += Step([]rune(s)).Hash()
	}
	return result
}

// FocusingPower returns the focusing power of the lens configuration.
func FocusingPower(input []string) int {
	numBoxes := 256
	boxes := make([]*Box, numBoxes)
	for i := 0; i < numBoxes; i++ {
		boxes[i] = NewBox(i)
	}

	steps := strings.Split(input[0], ",")
	for _, step := range steps {
		reg := regexp.MustCompile(`([a-z]+)([=-])(\d*)`)
		matches := reg.FindStringSubmatch(step)
		hash := Step([]rune(matches[1])).Hash()
		box := boxes[hash]

		if matches[2] == "=" {
			box.Add(matches[1], strconv.MustAtoi(matches[3]))
		} else {
			box.Remove(matches[1])
		}
	}

	result := 0
	for _, box := range boxes {
		result += box.FocusingPower()
	}
	return result
}

// Step represents a step in the input.
type Step []rune

// Hash returns a hash of the step.
func (s Step) Hash() int {
	result := 0
	for _, r := range s {
		result += int(r)
		result *= 17
		result %= 256
	}

	return result
}

// Box represents a box.
type Box struct {
	id           int
	slots        []int
	labelToSlots map[string]int
}

// NewBox returns a new box.
func NewBox(id int) *Box {
	return &Box{id: id, labelToSlots: make(map[string]int)}
}

// Remove removes the lens with the label from the box.
func (b *Box) Remove(label string) {
	slot, ok := b.labelToSlots[label]
	if !ok {
		return
	}

	b.slots[slot] = 0
	delete(b.labelToSlots, label)
}

// Add adds the lens with the label to the box.
func (b *Box) Add(label string, focalLength int) {
	slot, ok := b.labelToSlots[label]
	if ok {
		b.slots[slot] = focalLength
		return
	}
	b.slots = append(b.slots, focalLength)
	b.labelToSlots[label] = len(b.slots) - 1
}

// FocusingPower returns the focusing power of the box.
func (b *Box) FocusingPower() int {
	result := 0
	slotNum := 1
	boxNum := b.id + 1
	for _, slot := range b.slots {
		if slot == 0 {
			continue
		}
		result += boxNum * slotNum * slot
		slotNum++
	}
	return result
}

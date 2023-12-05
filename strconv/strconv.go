package strconv

import (
	"fmt"
	"strconv"
)

// MustAtoi converts the given string to an int or panics.
func MustAtoi(input string) int {
	n, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("could not convert %s to int", input))
	}
	return n
}

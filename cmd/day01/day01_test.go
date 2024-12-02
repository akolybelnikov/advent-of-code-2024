package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assertions := assert.New(t)
	input := `3   4
4   3
2   5
1   3
3   9
3   3
`

	t.Run("part 1", func(t *testing.T) {
		expected := 11
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 31
		actual := part2(input)

		assertions.Equal(actual, expected)
	})
}

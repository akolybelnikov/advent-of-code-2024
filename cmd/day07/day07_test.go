package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay07(t *testing.T) {
	assertions := assert.New(t)
	input := `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	t.Run("part 1", func(t *testing.T) {
		expected := int64(3749)
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := int64(11387)
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

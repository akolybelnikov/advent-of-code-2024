package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay04(t *testing.T) {
	assertions := assert.New(t)
	input := `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	t.Run("part 1", func(t *testing.T) {
		expected := 18
		actual := part1(input)

		assertions.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 9
		actual := part2(input)

		assertions.Equal(actual, expected)
	})
}

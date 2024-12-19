package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay19(t *testing.T) {
	assertions := assert.New(t)
	input := `
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

	t.Run("part 1", func(t *testing.T) {
		expected := 6
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 16
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

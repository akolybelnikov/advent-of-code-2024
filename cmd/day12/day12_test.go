package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12(t *testing.T) {
	assertions := assert.New(t)
	input := `
RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	t.Run("part 1", func(t *testing.T) {
		expected := 1930
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 1206
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay06(t *testing.T) {
	assertions := assert.New(t)
	input := `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	t.Run("part 1", func(t *testing.T) {
		expected := 41
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 6
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

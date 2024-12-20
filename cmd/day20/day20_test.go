package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay20(t *testing.T) {
	assertions := assert.New(t)
	input := `
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

	t.Run("part 1", func(t *testing.T) {
		expected := 44
		actual := part1(input, 2)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 285
		actual := part2(input, 50)

		assertions.Equal(expected, actual)
	})
}

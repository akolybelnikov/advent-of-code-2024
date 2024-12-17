package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16(t *testing.T) {
	assertions := assert.New(t)
	input := `
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

	input2 := `
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

	t.Run("part 1", func(t *testing.T) {
		expected := 7036
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 1 input 2", func(t *testing.T) {
		expected := 11048
		actual := part1(input2)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 45
		actual := part2(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2 input 2", func(t *testing.T) {
		expected := 64
		actual := part2(input2)

		assertions.Equal(expected, actual)
	})
}

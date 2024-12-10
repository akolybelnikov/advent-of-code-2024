package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	assertions := assert.New(t)
	input := `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	t.Run("part 1", func(t *testing.T) {
		expected := 36
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 81
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay09(t *testing.T) {
	assertions := assert.New(t)
	input := "2333133121414131402"

	t.Run("part 1", func(t *testing.T) {
		expected := 1928
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 2858
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

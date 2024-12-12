package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11(t *testing.T) {
	assertions := assert.New(t)
	input := "125 17"

	t.Run("part 1", func(t *testing.T) {
		expected := 55312
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 65601038650482
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

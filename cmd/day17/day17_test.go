package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay17(t *testing.T) {
	assertions := assert.New(t)
	input := ""

	t.Run("part 1", func(t *testing.T) {
		expected := 0
		actual := part1(input)

		assertions.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assertions.Equal(actual, expected)
	})
}

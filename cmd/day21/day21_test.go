package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay21(t *testing.T) {
	assertions := assert.New(t)
	input := `
029A
980A
179A
456A
379A`

	t.Run("part 1", func(t *testing.T) {
		var expected uint64 = 126384
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		var expected uint64 = 154115708116294
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

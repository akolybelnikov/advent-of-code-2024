package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay02(t *testing.T) {
	assertions := assert.New(t)
	input := `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 `

	t.Run("part 1", func(t *testing.T) {
		expected := 2
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 4
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

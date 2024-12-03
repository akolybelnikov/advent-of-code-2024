package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay03(t *testing.T) {
	assertions := assert.New(t)
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	t.Run("part 1", func(t *testing.T) {
		expected := 161
		actual := part1(input)

		assertions.Equal(actual, expected)
	})

	input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	t.Run("part 2", func(t *testing.T) {
		expected := 48
		actual := part2(input)

		assertions.Equal(actual, expected)
	})
}

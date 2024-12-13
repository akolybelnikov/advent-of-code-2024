package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13(t *testing.T) {
	assertions := assert.New(t)
	input := `
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

	t.Run("part 1", func(t *testing.T) {
		expected := 480
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 875318608908
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

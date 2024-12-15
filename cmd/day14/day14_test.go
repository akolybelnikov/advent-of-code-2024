package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay14(t *testing.T) {
	assertions := assert.New(t)
	input := `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	t.Run("part 1", func(t *testing.T) {
		expected := 12
		actual := part1(input, 11, 7)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

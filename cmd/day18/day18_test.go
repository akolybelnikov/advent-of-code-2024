package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//  ...#...
//  ..#..#.
//  ....#..
//  ...#..#
//  ..#..#.
//  .#..#..
//  #.#....

func TestDay18(t *testing.T) {
	assertions := assert.New(t)
	input := `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

	t.Run("part 1", func(t *testing.T) {
		expected := 22
		actual := part1(input, 6, 12)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := "6,1"
		actual := part2(input, 6, 12)

		assertions.Equal(expected, actual)
	})
}

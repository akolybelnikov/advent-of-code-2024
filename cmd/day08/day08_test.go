package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay08(t *testing.T) {
	assertions := assert.New(t)
	input := `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	t.Run("part 1", func(t *testing.T) {
		expected := 14
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 34
		actual := part2(input)

		assertions.Equal(expected, actual)
	})
}

func Test_extendVectors(t *testing.T) {
	type args struct {
		points []point
		bounds *map[point]bool
	}
	tests := []struct {
		name string
		args args
		want []point
	}{
		{
			name: "1",
			args: args{
				points: []point{{0, 6}, {3, 2}, {5, 6}},
				bounds: &map[point]bool{{10, 6}: true, {7, 10}: true},
			},
			want: []point{{x: 10, y: 6}, {x: 7, y: 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, extendVectors(tt.args.points, tt.args.bounds), "extendVectors(%v)", tt.args.points)
		})
	}
}

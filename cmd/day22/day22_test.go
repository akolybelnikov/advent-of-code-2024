package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay22(t *testing.T) {
	assertions := assert.New(t)
	input := `
1
10
100
2024`

	t.Run("part 1", func(t *testing.T) {
		var expected uint = 37327623
		actual := part1(input)

		assertions.Equal(expected, actual)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assertions.Equal(actual, expected)
	})
}

func TestEvolve(t *testing.T) {
	tests := []struct {
		name     string
		input    uint
		expected uint
	}{
		{
			name:     "123",
			input:    123,
			expected: 15887950,
		},
		{
			name:     "15887950",
			input:    15887950,
			expected: 16495136,
		},
		{
			name:     "16495136",
			input:    16495136,
			expected: 527345,
		},
		{
			name:     "527345",
			input:    527345,
			expected: 704524,
		},
		{
			name:     "704524",
			input:    704524,
			expected: 1553684,
		},
		{
			name:     "1553684",
			input:    1553684,
			expected: 12683156,
		},
		{
			name:     "12683156",
			input:    12683156,
			expected: 11100544,
		},
		{
			name:     "11100544",
			input:    11100544,
			expected: 12249484,
		},
		{
			name:     "12249484",
			input:    12249484,
			expected: 7753432,
		},
		{
			name:     "7753432",
			input:    7753432,
			expected: 5908254,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := evolve(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestEvolve2000(t *testing.T) {
	tests := []struct {
		name     string
		input    uint
		expected uint
	}{
		{
			name:     "1",
			input:    1,
			expected: 8685429,
		},
		{
			name:     "10",
			input:    10,
			expected: 4700978,
		},
		{
			name:     "100",
			input:    100,
			expected: 15273692,
		},
		{
			name:     "2024",
			input:    2024,
			expected: 8667524,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := evolve2000(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

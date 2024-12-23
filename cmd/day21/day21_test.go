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
		expected := 126384
		actual := part1(input)

		assertions.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assertions.Equal(actual, expected)
	})
}

func Test_Sequence_029A(t *testing.T) {
	tests := []struct {
		name           string
		initialKey     rune
		nextKey        rune
		expectedOutput string
	}{
		{
			name:           "A -> 0",
			initialKey:     'A',
			nextKey:        '0',
			expectedOutput: "<vA<AA>>^AvAA<^A>A",
		},
		{
			name:           "0 -> 2",
			initialKey:     '0',
			nextKey:        '2',
			expectedOutput: "v<<A>>^AvA^A",
		},
		{
			name:           "2  -> 9",
			initialKey:     '2',
			nextKey:        '9',
			expectedOutput: "<vA>^Av<<A>^A>AAvA^A",
		},
		{
			name:           "9 -> A",
			initialKey:     '9',
			nextKey:        'A',
			expectedOutput: "v<<A>A>^AAAvA<^A>A",
		},
		{
			name:           "edge_case_same_keys",
			initialKey:     '2',
			nextKey:        '2',
			expectedOutput: "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := newArea()
			result := a.sequence(tt.initialKey, tt.nextKey)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func Test_Sequence379A(t *testing.T) {
	tests := []struct {
		name           string
		initialKey     rune
		nextKey        rune
		expectedOutput string
	}{
		{
			name:           "A -> 3",
			initialKey:     'A',
			nextKey:        '3',
			expectedOutput: "v<<A>>^AvA^A",
		},
		{
			name:           "3 -> 7",
			initialKey:     '3',
			nextKey:        '7',
			expectedOutput: "<vA<AA>>^AAvA<^A>AAvA^A",
		},
		{
			name:           "7 -> 9",
			initialKey:     '7',
			nextKey:        '9',
			expectedOutput: "<vA>^AA<A>A",
		},
		{
			name:           "9 -> A",
			initialKey:     '9',
			nextKey:        'A',
			expectedOutput: "v<<A>A>^AAAvA<^A>A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := newArea()
			result := a.sequence(tt.initialKey, tt.nextKey)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

// A -> 3
// ^A
// <A>A
// <v<A>>^AvA^A

// 3 -> 7
// <<^^A
// v<<AA>^AA>A
// <vA<AA>>^AAvA<^A>AAvA^A

// 7 -> 9
// >>A
// vAA^A
// <vA>^AA<A>A

// 9 -> A
// vvvA
// <vAAA>^A
// <v<A>A>^AAAvA<^A>A

//
//+---+---+---+
//| 7 | 8 | 9 |
//+---+---+---+
//| 4 | 5 | 6 |
//+---+---+---+
//| 1 | 2 | 3 |
//+---+---+---+
//    | 0 | A |
//    +---+---+

//    +---+---+
//    | ^ | A |
//+---+---+---+
//| < | v | > |
//+---+---+---+
//
//<A                 | ^A | >^^A | vvvA
//v<<A >>^A           | <A>A | vA<^AA>A | <vAAA>^A
//<vA<AA>>^A vAA<^A>A | <v<A>>^AvA^A | <vA>^A<v<A>^A>AAvA^ | <v<A>A>^AAAvA<^A>A
//<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A

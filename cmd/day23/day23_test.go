package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay23(t *testing.T) {
	assertions := assert.New(t)
	input := `
kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

	t.Run("part 1", func(t *testing.T) {
		expected := 7
		actual := part1(input)

		assertions.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := "co,de,ka,ta"
		actual := part2(input)

		assertions.Equal(actual, expected)
	})
}

func TestFindSubNetworks(t *testing.T) {
	assertions := assert.New(t)

	tests := []struct {
		name     string
		input    network
		expected map[string]struct{}
	}{
		{
			name:     "empty network",
			input:    network{},
			expected: map[string]struct{}{},
		},
		{
			name: "no subnetworks",
			input: network{
				"a": {"b", "c"},
				"d": {"e"},
			},
			expected: map[string]struct{}{},
		},
		{
			name: "single subnetwork",
			input: network{
				"a": {"b", "c"},
				"b": {"a", "c"},
				"c": {"a", "b"},
			},
			expected: map[string]struct{}{
				"a,b,c": {},
			},
		},
		{
			name: "single subnetworks multiple times",
			input: network{
				"a": {"b", "c", "d"},
				"b": {"a", "c"},
				"c": {"a", "b"},
				"d": {"a", "e"},
				"e": {"d"},
			},
			expected: map[string]struct{}{
				"a,b,c": {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.input.findSubnetworks()

			assertions.Equal(tt.expected, actual)
		})
	}
}

package main

import (
	"fmt"
	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
	"image"
	"math"
	"os"
	"sort"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day21.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) uint64 {
	return run(input, 2)
}

// part two
func part2(input string) uint64 {
	return run(input, 25)
}

// run processes a given input string and returns the final aggregated result as a uint64 value.
func run(input string, maxDepth int) uint64 {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	var result uint64
	for _, code := range lines {
		length := FindMinLengthWrapper(code, maxDepth)
		result += length * uint64(Number(code))
	}
	return result
}

// KeyBoard represents one of the keypads.
type KeyBoard struct {
	Keys    map[rune]image.Point
	KeysRev map[image.Point]rune
	Pos     image.Point
	PPos    image.Point
}

// NewKeyBoard initializes a keypad with the given rows, columns, and layout.
func NewKeyBoard(rows, cols int, layout string) *KeyBoard {
	kb := &KeyBoard{
		Keys:    make(map[rune]image.Point),
		KeysRev: make(map[image.Point]rune),
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			key := rune(layout[cols*r+c])
			point := image.Point{Y: r, X: c}
			kb.Keys[key] = point
			kb.KeysRev[point] = key
			if key == 'A' {
				kb.Pos = point
			}
		}
	}

	kb.PPos = kb.Pos
	return kb
}

// IsValidPath checks if the given path is valid.
// IsValidPath checks if the given path is valid.
func (kb *KeyBoard) IsValidPath(moves string) bool {
	pos := kb.Pos
	for _, m := range moves {
		switch m {
		case '<':
			pos = pos.Add(image.Point{X: -1, Y: 0})
		case '>':
			pos = pos.Add(image.Point{X: 1, Y: 0})
		case '^':
			pos = pos.Add(image.Point{X: 0, Y: -1})
		case 'v':
			pos = pos.Add(image.Point{X: 0, Y: 1})
		default:
			continue
		}
		if kb.KeysRev[pos] == '*' {
			return false
		}
	}
	return true
}

// MoveTo updates the current position of the keypad to the given key.
func (kb *KeyBoard) MoveTo(key rune) {
	kb.Pos = kb.Keys[key]
}

// GetPaths returns all valid paths from the current position to the given key.
// GetPaths returns all valid paths from the current position to the given key.
func (kb *KeyBoard) GetPaths(key rune) []string {
	target := kb.Keys[key]
	diff := target.Sub(kb.Pos)

	// Generate a "Manhattan path".
	path := ""
	for i := 0; i < int(math.Abs(float64(diff.X))); i++ {
		if diff.X < 0 {
			path += "<"
		} else {
			path += ">"
		}
	}

	for i := 0; i < int(math.Abs(float64(diff.Y))); i++ {
		if diff.Y < 0 {
			path += "^"
		} else {
			path += "v"
		}
	}

	// Permute paths to get all possibilities.
	var paths []string
	runes := []rune(path)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

	for {
		p := string(runes) + "A"
		if kb.IsValidPath(p) {
			paths = append(paths, p)
		}
		if !nextPermutation(runes) {
			break
		}
	}

	return paths
}

// Helper function to generate the next lexicographical permutation of a slice.
func nextPermutation(runes []rune) bool {
	n := len(runes)
	i := n - 2
	for i >= 0 && runes[i] >= runes[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for runes[j] <= runes[i] {
		j--
	}
	runes[i], runes[j] = runes[j], runes[i]
	reverse(runes[i+1:])
	return true
}

func reverse(runes []rune) {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

// Number converts a code into an integer by ignoring 'A' and forming the number.
func Number(code string) int {
	result := 0
	for _, c := range code {
		if c == 'A' {
			continue
		}
		result = result*10 + int(c-'0')
	}
	return result
}

// FindMinLength recursively evaluates sequences to find the shortest path length.
func FindMinLength(sequence string, depth, maxDepth int, cache map[string]uint64) uint64 {
	// Terminate the recursion.
	if depth == maxDepth {
		return uint64(len(sequence))
	}

	// Check the cache.
	cacheKey := fmt.Sprintf("%d:%s", depth, sequence)
	if val, exists := cache[cacheKey]; exists {
		return val
	}

	dir := NewKeyBoard(2, 3, "*^A<v>")
	var totalMinLength uint64

	for _, c := range sequence {
		paths := dir.GetPaths(c)
		dir.MoveTo(c)

		minLength := uint64(math.MaxUint64)
		for _, path := range paths {
			length := FindMinLength(path, depth+1, maxDepth, cache)
			if length < minLength {
				minLength = length
			}
		}
		totalMinLength += minLength
	}

	cache[cacheKey] = totalMinLength
	return totalMinLength
}

// FindMinLengthWrapper serves as a wrapper for the find_min_length function.
func FindMinLengthWrapper(code string, maxDepth int) uint64 {
	pad := NewKeyBoard(4, 3, "789456123*0A")

	allPaths := []string{""}
	for _, c := range code {
		paths := pad.GetPaths(c)
		pad.MoveTo(c)

		var newPaths []string
		for _, r := range allPaths {
			for _, p := range paths {
				newPaths = append(newPaths, r+p)
			}
		}
		allPaths = newPaths
	}

	cache := make(map[string]uint64)
	minLength := uint64(math.MaxUint64)
	for _, path := range allPaths {
		length := FindMinLength(path, 0, maxDepth, cache)
		if length < minLength {
			minLength = length
		}
	}

	return minLength
}

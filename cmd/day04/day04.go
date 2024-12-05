package main

import (
	"fmt"
	"os"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

type Direction int

const (
	Up Direction = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

const (
	X rune = 88
	M rune = 77
	A rune = 65
	S rune = 83
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day04.txt")
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

func findPaths(x, y, size int) []Direction {
	switch {
	case x < 3 && y < 3:
		return []Direction{Right, DownRight, Down}
	case x >= 3 && x < size-3 && y < 3:
		return []Direction{Right, DownRight, Down, Left, DownLeft}
	case x >= size-3 && y < 3:
		return []Direction{Down, DownLeft, Left}
	case x < 3 && y < size-3:
		return []Direction{Up, UpRight, Right, DownRight, Down}
	case x >= 3 && x < size-3 && y < size-3:
		return []Direction{Up, UpRight, Right, DownRight, Down, DownLeft, Left, UpLeft}
	case x >= size-3 && y < size-3:
		return []Direction{Up, Down, DownLeft, Left, UpLeft}
	case x < 3 && y >= size-3:
		return []Direction{Up, UpRight, Right}
	case x >= 3 && x < size-3 && y >= size-3:
		return []Direction{Up, UpRight, Right, Left, UpLeft}
	case x >= size-3 && y >= size-3:
		return []Direction{Up, Left, UpLeft}
	default:
		return []Direction{}
	}
}

func findXMASRunes(x, y int, direction Direction, matrix *[][]rune) int {
	runes := []rune{M, A, S}
	for _, r := range runes {
		switch direction {
		case Up:
			y--
		case UpRight:
			x++
			y--
		case Right:
			x++
		case DownRight:
			x++
			y++
		case Down:
			y++
		case DownLeft:
			x--
			y++
		case Left:
			x--
		case UpLeft:
			x--
			y--
		}
		if (*matrix)[x][y] != r {
			return 0
		}
	}

	return 1
}

func checkMASToken(x, y int, matrix *[][]rune) int {
	ul := (*matrix)[x-1][y-1]
	ur := (*matrix)[x+1][y-1]
	dr := (*matrix)[x+1][y+1]
	dl := (*matrix)[x-1][y+1]

	switch {
	case ul == M && ur == M && dr == S && dl == S:
		return 1
	case ul == M && ur == S && dr == S && dl == M:
		return 1
	case ul == S && ur == M && dr == M && dl == S:
		return 1
	case ul == S && ur == S && dr == M && dl == M:
		return 1
	default:
		return 0

	}
}

// part one
func part1(input string) int {
	var sum int
	lines, err := utils.ParseLines(input)
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}
	matrix := utils.ConvertLinesToRuneSlices(lines)

	for i, runes := range matrix {
		for j, r := range runes {
			if r != X {
				continue
			}
			paths := findPaths(i, j, len(matrix))
			for _, path := range paths {
				sum += findXMASRunes(i, j, path, &matrix)
			}
		}
	}

	return sum
}

// part two
func part2(input string) int {
	var sum int
	lines, err := utils.ParseLines(input)
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}
	matrix := utils.ConvertLinesToRuneSlices(lines)

	for i, runes := range matrix {
		for j, r := range runes {
			if r == A && i > 0 && j > 0 && i < len(matrix)-1 && j < len(matrix[0])-1 {
				sum += checkMASToken(i, j, &matrix)
			}
		}
	}

	return sum
}

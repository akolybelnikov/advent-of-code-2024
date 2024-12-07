package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

type coordinate [2]int
type direction int

const (
	CROSS    rune = '+'
	EMPTY    rune = '.'
	GUARD    rune = '^'
	OBSTACLE rune = '#'
	hPIPE    rune = '-'
	vPIPE    rune = '|'
)

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day06.txt")
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

func createGridWithGuard(input string) ([][]rune, coordinate) {
	lines, err := utils.ParseLines(input)
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}

	matrix := utils.ConvertLinesToRuneSlices(lines)

	var guardCoordinates coordinate
	for i, row := range matrix {
		for j, cell := range row {
			if cell == GUARD {
				guardCoordinates = coordinate{i, j}
				break
			}
		}
	}

	return matrix, guardCoordinates
}

func nextPosition(cur coordinate, direction direction) coordinate {
	switch direction {
	case UP:
		return coordinate{cur[0] - 1, cur[1]}
	case RIGHT:
		return coordinate{cur[0], cur[1] + 1}
	case DOWN:
		return coordinate{cur[0] + 1, cur[1]}
	case LEFT:
		return coordinate{cur[0], cur[1] - 1}
	default:
		return coordinate{-1, -1}
	}
}

// part one
func part1(input string) int {
	grid, cur := createGridWithGuard(input)
	uniqueCoordinates := make(map[coordinate]int)
	currentDirection := UP

	for {
		uniqueCoordinates[cur]++
		next := nextPosition(cur, currentDirection)
		if next[0] < 0 || next[1] < 0 || next[0] > len(grid)-1 || next[1] > len(grid[0])-1 {
			break
		}
		if grid[next[0]][next[1]] == OBSTACLE {
			currentDirection = (currentDirection + 1) % 4
		} else {
			cur = next
		}
	}

	return len(uniqueCoordinates)
}

func hashCell(cell coordinate, dir direction) string {
	hashGenerator := sha256.New()
	hashGenerator.Write([]byte(string(rune(cell[0]))))
	hashGenerator.Write([]byte(string(rune(cell[1]))))
	hashGenerator.Write([]byte(string(rune(dir))))

	return hex.EncodeToString(hashGenerator.Sum(nil))
}

func mutateCell(pos coordinate, matrix *[][]rune, dir direction) {
	cur := (*matrix)[pos[0]][pos[1]]
	switch {
	case (cur == hPIPE && (dir == UP || dir == DOWN)) || (cur == vPIPE && (dir == LEFT || dir == RIGHT)):
		(*matrix)[pos[0]][pos[1]] = CROSS
	case cur == GUARD:
		(*matrix)[pos[0]][pos[1]] = vPIPE
	case cur == EMPTY && (dir == UP || dir == DOWN):
		(*matrix)[pos[0]][pos[1]] = vPIPE
	case cur == EMPTY && (dir == LEFT || dir == RIGHT):
		(*matrix)[pos[0]][pos[1]] = hPIPE
	}
}

// part two
func part2(input string) int {
	grid, cur := createGridWithGuard(input)
	currentDirection := UP
	path := make(map[coordinate]int)

	for {
		next := nextPosition(cur, currentDirection)

		if isOutOfBounds(next, len(grid), len(grid[0])) {
			break
		}

		if grid[next[0]][next[1]] == OBSTACLE {
			currentDirection = (currentDirection + 1) % 4
			grid[cur[0]][cur[1]] = CROSS
		} else {
			mutateCell(cur, &grid, currentDirection)
			cur = next
			path[cur]++
		}
	}

	return countLoops(input, path)
}

func countLoops(input string, path map[coordinate]int) int {
	loops := 0

	for pos := range path {
		grid, cursor := createGridWithGuard(input)
		grid[pos[0]][pos[1]] = OBSTACLE
		currentDirection := UP
		cellVisits := make(map[string]int)

		for {
			next := nextPosition(cursor, currentDirection)

			if isOutOfBounds(next, len(grid), len(grid[0])) {
				break
			}

			if grid[next[0]][next[1]] == OBSTACLE {
				currentDirection = (currentDirection + 1) % 4
				grid[cursor[0]][cursor[1]] = CROSS
			} else {
				mutateCell(cursor, &grid, currentDirection)
				cellHash := hashCell(cursor, currentDirection)
				visitCount, exists := cellVisits[cellHash]

				if exists && visitCount > 2 {
					loops++
					break
				}

				cellVisits[cellHash]++
				cursor = next
			}
		}
	}

	return loops
}

func isOutOfBounds(point coordinate, rows int, cols int) bool {
	return point[0] < 0 || point[1] < 0 || point[0] >= rows || point[1] >= cols
}

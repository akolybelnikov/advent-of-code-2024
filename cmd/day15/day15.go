package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day15.txt")
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
func part1(input string) int {
	blocks, err := utils.ParseBlocksOfLines(input)
	utils.HandleErr(err)

	instructions := sanitize(blocks[1])
	grid, robot := createGrid(blocks[0])
	directions := map[rune]image.Point{
		'<': {-1, 0},
		'>': {1, 0},
		'^': {0, -1},
		'v': {0, 1},
	}

	visualize(&grid)

	for _, dir := range instructions {
		move(&grid, &robot, directions[dir])
	}

	visualize(&grid)

	return sumOfBoxesGPS(&grid)
}

// part two
func part2(input string) int {
	blocks, err := utils.ParseBlocksOfLines(input)
	utils.HandleErr(err)

	instructions := sanitize(blocks[1])
	expanded := expandInput(blocks[0])
	grid, robot := createGrid(expanded)

	directions := map[rune]image.Point{
		'<': {-1, 0},
		'>': {1, 0},
		'^': {0, -1},
		'v': {0, 1},
	}

	visualize(&grid)

	for _, dir := range instructions {
		move(&grid, &robot, directions[dir])
	}

	visualize(&grid)

	return 0
}

// sanitize is used to trim and join the instructions
func sanitize(input []string) string {
	var builder strings.Builder
	for _, s := range input {
		s = strings.Trim(s, "\n\r ")
		builder.WriteString(s)
	}

	return builder.String()
}

// createGrid converts the puzzle input string into a map with image.Point as keys and runes as values
func createGrid(input []string) (map[image.Point]rune, image.Point) {
	grid := make(map[image.Point]rune)
	var robot image.Point

	for y, s := range input {
		for x, r := range s {
			grid[image.Point{X: x, Y: y}] = r
			if r == '@' {
				robot = image.Point{X: x, Y: y}
			}
		}
	}

	return grid, robot
}

// move mutates the grid and the robot coordinates for each instruction
func move(grid *map[image.Point]rune, robot *image.Point, dir image.Point) {
	next := robot.Add(dir)
	switch (*grid)[next] {
	case '.':
		(*grid)[next] = '@'
		(*grid)[*robot] = '.'
		*robot = next
	case 'O':
		num := lookup(grid, next, dir)
		if num > 0 {
			(*grid)[next] = '@'
			(*grid)[*robot] = '.'
			*robot = next
			for i := 1; i <= num; i++ {
				cell := next.Add(dir.Mul(i))
				(*grid)[cell] = 'O'
			}
		}
	case '[', ']':
		if dir.X == 0 && (dir.Y == -1 || dir.Y == 1) {
			num := lookup(grid, next.Add(dir), dir)
			if num > 0 {
				(*grid)[next] = '@'
				(*grid)[*robot] = '.'
				*robot = next
				for i := 1; i <= num; i++ {
					cell := next.Add(dir.Mul(i))
					if (i%2 != 0 && dir.Y == 1) || (i%2 == 0 && dir.Y == -1) {
						(*grid)[cell] = '['
					} else {
						(*grid)[cell] = ']'
					}
				}
			}
		} else {

		}
	default:

	}
}

// lookup counts the number of moves if robot is facing a box
func lookup(grid *map[image.Point]rune, box image.Point, dir image.Point) int {
	num := 0
	next := box.Add(dir)
	switch (*grid)[next] {
	case '.':
		num += 1
	case 'O':
		num += 1
		for (*grid)[next] != '.' && (*grid)[next] != '#' {
			num += 1
			next = next.Add(dir)
		}
		if (*grid)[next] == '#' {
			return 0
		}
	case '[', ']':
		num += 2
		for (*grid)[next] != '.' && (*grid)[next] != '#' {
			num += 2
			next = next.Add(dir.Mul(2))
		}
		if (*grid)[next] == '#' {
			return 0
		}
	default:
		return 0
	}

	return num
}

func verticalLookup(grid *map[image.Point]rune, side image.Point, dir image.Point) *[]image.Point {
	boxes := completeBox((*grid)[side], side)
	nextLeft, nextRight := boxes[0].Add(dir), boxes[1].Add(dir)
	switch {
	case (*grid)[nextLeft] == '.' && (*grid)[nextRight] == '.':
		return &boxes
	case (*grid)[nextLeft] == '#' || (*grid)[nextRight] == '#':
		return nil
	case (*grid)[nextLeft] == '[', (*grid)[nextLeft] == ']':

	}

	return &boxes
}

func completeBox(side rune, point image.Point) []image.Point {
	if side == '[' {
		return []image.Point{point, point.Add(image.Point{X: 1, Y: 0})}
	} else {
		return []image.Point{point.Add(image.Point{X: -1, Y: 0}), point}
	}
}

// sumOfBoxesGPS calculates the sum according to the puzzle rules
func sumOfBoxesGPS(grid *map[image.Point]rune) int {
	sum := 0

	for p, r := range *grid {
		if r == 'O' {
			sum += p.Y*100 + p.X
		}
	}

	return sum
}

// expandInput doubles the width of the original grid according to the puzzle rules for part 2
func expandInput(input []string) []string {
	builder := strings.Builder{}
	res := make([]string, len(input))
	for i, s := range input {
		for _, r := range s {
			switch r {
			case '#':
				builder.WriteString("##")
			case 'O':
				builder.WriteString("[]")
			case '@':
				builder.WriteString("@.")
			default:
				builder.WriteString("..")
			}
		}
		res[i] = builder.String()
		builder.Reset()
	}

	return res
}

// visualize converts the map of image.Points into their string representations to print it out
func visualize(grid *map[image.Point]rune) {
	// Determine the grid bounds
	minX, minY := 0, 0
	maxX, maxY := 0, 0
	for p := range *grid {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	// Build the grid as a string
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if r, ok := (*grid)[image.Point{X: x, Y: y}]; ok {
				fmt.Print(string(r))
			} else {
				fmt.Print(".") // Default empty space
			}
		}
		fmt.Println()
	}
}

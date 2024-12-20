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

	for _, dir := range instructions {
		move(&grid, &robot, directions[dir])
	}

	return sumOfBoxesGPS(&grid)
}

// part two
func part2(input string) int {
	blocks := strings.Split(strings.TrimSpace(input), "\n\r")

	r := strings.NewReplacer("#", "##", "O", "[]", ".", "..", "@", "@.")

	return run(r.Replace(blocks[0]), blocks[1])
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
	default:
		return 0
	}

	return num
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

func run(input, moves string) int {
	grid, robot := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Fields(input) {
		for x, r := range s {
			if r == '@' {
				robot = image.Point{X: x, Y: y}
				r = '.'
			}
			grid[image.Point{X: x, Y: y}] = r
		}
	}

	delta := map[rune]image.Point{
		'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0},
		'[': {1, 0}, ']': {-1, 0},
	}

loop:
	for _, r := range strings.ReplaceAll(moves, "\n", "") {
		queue, boxes := []image.Point{robot}, map[image.Point]rune{}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if _, ok := boxes[p]; ok {
				continue
			}
			boxes[p] = grid[p]

			switch n := p.Add(delta[r]); grid[n] {
			case '#':
				continue loop
			case '[', ']':
				queue = append(queue, n.Add(delta[grid[n]]))
				fallthrough
			case 'O':
				queue = append(queue, n)
			}
		}

		for b := range boxes {
			grid[b] = '.'
		}
		for b := range boxes {
			grid[b.Add(delta[r])] = boxes[b]
		}
		robot = robot.Add(delta[r])
	}

	gps := 0
	for p, r := range grid {
		if r == 'O' || r == '[' {
			gps += 100*p.Y + p.X
		}
	}
	return gps
}

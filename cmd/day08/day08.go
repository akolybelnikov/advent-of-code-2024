package main

import (
	"fmt"
	"os"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

type grid [][]rune

type point struct {
	x, y int
}

func (p *point) diff(q point) point {
	return point{p.x - q.x, p.y - q.y}
}

func (p *point) add(q point) point {
	return point{p.x + q.x, p.y + q.y}
}

func (g *grid) inspect() (map[rune][]point, map[point]bool) {
	m := make(map[rune][]point)
	bounds := make(map[point]bool)

	for x, row := range *g {
		for y, cell := range row {
			pos := point{x, y}
			bounds[pos] = true
			if cell != '.' {
				m[cell] = append(m[cell], pos)
			}
		}
	}

	return m, bounds
}

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day08.txt")
	utils.HandleErr(err)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	uniquePositions := make(map[point]struct{})

	rl, err := utils.ParseRuneLines(input)
	utils.HandleErr(err)

	g := grid(rl)
	m, bounds := g.inspect()

	for _, positions := range m {
		for _, p := range extendVectors(positions, &bounds) {
			uniquePositions[p] = struct{}{}
		}
	}

	return len(uniquePositions)
}

// part two
func part2(input string) int {
	uniquePositions := make(map[point]struct{})

	rl, err := utils.ParseRuneLines(input)
	utils.HandleErr(err)

	g := grid(rl)
	m, bounds := g.inspect()

	for _, positions := range m {
		if len(positions) > 1 {
			for _, pos := range positions {
				uniquePositions[pos] = struct{}{}
			}
			for _, pos := range extendVectors2(positions, &bounds) {
				uniquePositions[pos] = struct{}{}
			}
		}
	}

	return len(uniquePositions)
}

func extendVectors(points []point, bounds *map[point]bool) []point {
	var res []point

	for _, p := range points {
		for _, q := range points {
			if p == q {
				continue
			}
			if r := p.add(p.diff(q)); (*bounds)[r] {
				res = append(res, r)
			}
		}
	}

	return res
}

func extendVectors2(points []point, bounds *map[point]bool) []point {
	var res []point

	for _, p := range points {
		for _, q := range points {
			if p == q {
				continue
			}
			for d := q.diff(p); (*bounds)[q]; q = q.add(d) {
				res = append(res, q)
			}
		}
	}

	return res
}

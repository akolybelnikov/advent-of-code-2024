package main

import (
	"fmt"
	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
	"image"
	"os"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day12.txt")
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
	grid, err := utils.ParseRuneLines(input)
	utils.HandleErr(err)

	res, _ := traversRegions(grid)

	return res
}

// part two
func part2(input string) int {
	grid, err := utils.ParseRuneLines(input)
	utils.HandleErr(err)

	_, res := traversRegions(grid)

	return res
}

func traversRegions(input [][]rune) (int, int) {
	grid := map[image.Point]rune{}
	for y, s := range input {
		for x, r := range s {
			grid[image.Point{X: x, Y: y}] = r
		}
	}

	seen := map[image.Point]bool{}
	res1, res2 := 0, 0
	for c := range grid {
		if seen[c] {
			continue
		}
		seen[c] = true

		area := 1
		perimeter, sides := 0, 0
		queue := []image.Point{c}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				if n := p.Add(d); grid[n] != grid[p] {
					perimeter++
					r := p.Add(image.Point{X: -d.Y, Y: d.X})
					if grid[r] != grid[p] || grid[r.Add(d)] == grid[p] {
						sides++
					}
				} else if !seen[n] {
					seen[n] = true
					queue = append(queue, n)
					area++
				}
			}
		}
		res1 += area * perimeter
		res2 += area * sides
	}
	return res1, res2
}

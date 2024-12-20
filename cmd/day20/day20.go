package main

import (
	"fmt"
	"image"
	"os"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

const (
	START = 'S'
	END   = 'E'
	WALL  = '#'
	EMPTY = '.'
)

type pathPoint struct {
	cheats []image.Point
	ps     int
}

type grid struct {
	start, end image.Point
	path       map[image.Point]*pathPoint
	points     map[image.Point]rune
	directions []image.Point
}

// visualize converts the map of image.Points into their string representations to print it out
func (g *grid) visualize() {
	// Determine the grid bounds
	minX, minY := 0, 0
	maxX, maxY := 0, 0
	for p := range g.points {
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
			if r, ok := g.points[image.Point{X: x, Y: y}]; ok {
				fmt.Print(string(r))
			}
		}
		fmt.Println()
	}
}

// traverse explores the grid from the start point to the end point by updating paths with picoseconds and
// discovering potential cheat passages on the way
func (g *grid) traverse() {
	next := g.start

	for next != g.end {
		p := next
		cur := g.path[p]
		// discover the 4 possible next moves
		for _, dir := range g.directions {
			np := p.Add(dir)
			if g.points[np] == WALL {
				// if next move is wall, let us lookup behind the wall. If next cell is empty, this could be a cheat!
				lookup := np.Add(dir)
				r, ok := g.points[lookup]
				// this lets us prevent marking cells on the path with lower picoseconds count as cheats
				_, pok := g.path[lookup]
				// if the looked up cell has not been discovered so far, it lies ahead on the path, and it is a cheat!
				if ok && (r == EMPTY || r == END) && !pok {
					cur.cheats = append(cur.cheats, lookup)
				}
			} else {
				// if next cell is empty, it is the next step in the path as we move strictly between the walls
				if _, ok := g.path[np]; !ok {
					g.path[np] = &pathPoint{ps: cur.ps + 1}
					next = np
				}
			}
		}
	}
}

func (g *grid) count(min int) int {
	ans := 0
	cheats := make(map[int]int)

	for _, pp := range g.path {
		//fmt.Println(pos, pp)
		for _, cheat := range pp.cheats {
			diff := g.path[cheat].ps - pp.ps - 2
			//fmt.Println(diff, cheat, g.path[cheat])
			cheats[diff]++
		}
	}

	for k, v := range cheats {
		if k >= min {
			ans += v
		}
	}

	return ans
}

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day20.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input, 100))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string, min int) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	g, err := newGrid(lines)
	utils.HandleErr(err)

	g.traverse()

	return g.count(min)
}

// part two
func part2(input string) int {
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	g, err := newGrid(lines)
	utils.HandleErr(err)

	g.traverse()

	return 0
}

func newGrid(lines []string) (grid, error) {
	g := grid{}
	g.points = map[image.Point]rune{}
	g.path = map[image.Point]*pathPoint{}
	for y, s := range lines {
		for x, r := range s {
			if r == START {
				g.start = image.Point{X: x, Y: y}
				g.path[g.start] = &pathPoint{ps: 0}
			}
			if r == END {
				g.end = image.Point{X: x, Y: y}
			}
			g.points[image.Point{X: x, Y: y}] = r
		}
	}
	g.directions = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	zero := image.Point{X: 0, Y: 0}
	if g.start == zero || g.end == zero {
		return g, fmt.Errorf("start or end point is not found")
	}

	return g, nil
}

package main

import (
	"container/heap"
	"fmt"
	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
	"image"
	"math"
	"os"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day18.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input, 70, 1024))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input, 70, 1024))

	os.Exit(0)
}

// part one
func part1(input string, size int, num int) int {
	parsed, err := utils.ParseLines(input)
	utils.HandleErr(err)

	grid := makeGrid(size+1, parsed[:num])
	cost, _ := traverse(grid, size)

	return cost
}

// part two
func part2(input string, size int, num int) string {
	parsed, err := utils.ParseLines(input)
	utils.HandleErr(err)

	grid := makeGrid(size+1, parsed[:num])
	for _, b := range parsed[num:] {
		p := image.Point{}
		_, _ = fmt.Sscanf(b, "%d,%d", &p.X, &p.Y)
		(*grid)[p] = '#'
		cost, _ := traverse(grid, size)
		if cost == -1 {
			return b
		}
	}

	return ""
}

func makeGrid(size int, bytes []string) *map[image.Point]rune {
	grid := map[image.Point]rune{}
	for y := range size {
		for x := range size {
			grid[image.Point{X: x, Y: y}] = '.'
		}
	}

	for _, b := range bytes {
		var i image.Point
		_, _ = fmt.Sscanf(b, "%d,%d", &i.X, &i.Y)
		grid[i] = '#'
	}

	return &grid
}

func printGrid(grid *map[image.Point]rune, size int) {
	for y := 0; y <= size; y++ {
		for x := 0; x <= size; x++ {
			if r, ok := (*grid)[image.Point{X: x, Y: y}]; ok {
				fmt.Print(string(r))
			} else {
				fmt.Print(".") // Default empty space
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

type QI struct {
	Pos  image.Point
	Cost int
}

// PriorityQueue implements a min-heap for QI
type PriorityQueue []QI

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Cost < (*pq)[j].Cost
}

func (pq *PriorityQueue) Swap(i, j int) {
	p := *pq
	p[i], p[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(QI))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func traverse(grid *map[image.Point]rune, size int) (int, int) {
	start := image.Point{X: 0, Y: 0}
	target := image.Point{X: size, Y: size}
	dist := map[image.Point]int{}
	for p := range *grid {
		dist[p] = math.MaxInt32 // Set "infinity" distance for all points
	}

	// Priority queue for processing nodes
	queue := &PriorityQueue{}
	heap.Init(queue)
	heap.Push(queue, QI{Pos: start, Cost: 0}) // Start position with cost 0
	dist[start] = 0

	// Directions for movement (up, right, down, left)
	directions := []image.Point{
		{0, -1}, // Up
		{1, 0},  // Right
		{0, 1},  // Down
		{-1, 0}, // Left
	}

	// Track visited points for path length calculation
	visited := map[image.Point]struct{}{}

	for queue.Len() > 0 {
		current := heap.Pop(queue).(QI)
		if _, seen := visited[current.Pos]; seen {
			continue
		}
		visited[current.Pos] = struct{}{}

		// If we reach the target, return the cost and path length
		if current.Pos == target {
			return current.Cost, len(visited)
		}

		// Check neighbors
		for _, d := range directions {
			next := current.Pos.Add(d)
			if cell, ok := (*grid)[next]; !ok || cell == '#' {
				continue // Skip out-of-bounds or obstacle cells
			}
			newCost := current.Cost + 1
			if newCost < dist[next] {
				dist[next] = newCost
				heap.Push(queue, QI{Pos: next, Cost: newCost})
			}
		}
	}

	// If target is unreachable, return -1
	return -1, len(visited)
}

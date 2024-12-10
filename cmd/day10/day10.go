package main

import (
	"container/heap"
	"fmt"
	"os"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

type node struct {
	x, y, weight, priority, index int
	path                          [][2]int
}

// queue implements heap.Interface
type queue []*node

func (pq *queue) Len() int { return len(*pq) }
func (pq *queue) Less(i, j int) bool {
	// Min-heap based on the priority (weight)
	return (*pq)[i].priority < (*pq)[j].priority
}
func (pq *queue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}
func (pq *queue) Push(x interface{}) {
	item := x.(*node)
	item.index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *queue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // Mark as removed
	*pq = old[0 : n-1]
	return item
}

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day10.txt")
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
	var sum int
	grid, err := utils.ParseIntLines(input)
	utils.HandleErr(err)

	heads, ends := findPathHeads(&grid)
	for _, head := range heads {
		paths := findAllPaths(grid, head, ends)
		for _, uniquePaths := range paths {
			if len(uniquePaths) > 0 {
				sum += 1
			}
		}
	}

	return sum
}

// part two
func part2(input string) int {
	var sum int
	grid, err := utils.ParseIntLines(input)
	utils.HandleErr(err)

	heads, ends := findPathHeads(&grid)
	for _, head := range heads {
		paths := findAllPaths(grid, head, ends)
		for _, uniquePaths := range paths {
			sum += len(uniquePaths)
		}
	}

	return sum
}

func findPathHeads(grid *[][]int) ([][2]int, [][2]int) {
	var heads [][2]int
	var ends [][2]int

	for x := 0; x < len(*grid); x++ {
		for y := 0; y < len((*grid)[0]); y++ {
			if (*grid)[x][y] == 0 {
				heads = append(heads, [2]int{x, y})
			}
			if (*grid)[x][y] == 9 {
				ends = append(ends, [2]int{x, y})
			}
		}
	}

	return heads, ends
}

// Pathfinding function to find all unique paths to each destination
func findAllPaths(grid [][]int, start [2]int, destinations [][2]int) map[[2]int][][][2]int {
	rows, cols := len(grid), len(grid[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, down, left, up

	// Priority Queue
	pq := &queue{}
	heap.Init(pq)

	// Destinations to track
	destSet := make(map[[2]int]bool)
	for _, dest := range destinations {
		destSet[dest] = true
	}

	// Map to store all paths for each destination
	allPaths := make(map[[2]int][][][2]int)

	// Push the start cell into the priority queue
	heap.Push(pq, &node{
		x:        start[0],
		y:        start[1],
		weight:   grid[start[0]][start[1]],
		priority: grid[start[0]][start[1]],
		path:     [][2]int{start},
	})

	for pq.Len() > 0 {
		// Pop the item with the smallest priority (weight)
		item := heap.Pop(pq).(*node)
		currentPos := [2]int{item.x, item.y}

		// If the current cell is a destination and the path is valid, record it
		if destSet[currentPos] && item.weight == 9 {
			// Ensure the path is unique before adding
			isUnique := true
			for _, existingPath := range allPaths[currentPos] {
				if equalPaths(existingPath, item.path) {
					isUnique = false
					break
				}
			}
			if isUnique {
				allPaths[currentPos] = append(allPaths[currentPos], item.path)
			}
			// Continue exploring to find all unique paths
		}

		// Explore neighbors
		for _, d := range directions {
			nx, ny := item.x+d[0], item.y+d[1]

			// Check bounds
			if nx >= 0 && ny >= 0 && nx < rows && ny < cols {
				nextWeight := grid[nx][ny]
				// Ensure weights increase by exactly 1
				if nextWeight == item.weight+1 {
					heap.Push(pq, &node{
						x:        nx,
						y:        ny,
						weight:   nextWeight,
						priority: nextWeight,
						path:     append(append([][2]int{}, item.path...), [2]int{nx, ny}),
					})
				}
			}
		}
	}

	return allPaths
}

// Helper function to check if two paths are equal
func equalPaths(path1, path2 [][2]int) bool {
	if len(path1) != len(path2) {
		return false
	}
	for i := range path1 {
		if path1[i] != path2[i] {
			return false
		}
	}
	return true
}

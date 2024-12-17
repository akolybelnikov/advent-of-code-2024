package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"image"
	"maps"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

type Position struct {
	x, y      int
	direction string
}

// State represents a position in the grid with a specific direction
type State struct {
	x, y      int
	direction string
	cost, h   int // Total cost (g) and heuristic (h)
	index     int // Index in the priority queue
}

// PriorityQueue implements a min-heap for A* states
type PriorityQueue []*State

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].cost+(*pq)[i].h < (*pq)[j].cost+(*pq)[j].h
}
func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	state := x.(*State)
	state.index = n
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	state := old[n-1]
	state.index = -1
	*pq = old[0 : n-1]
	return state
}

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day16.txt")
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

	startX, startY := 1, len(grid)-2
	goalX, goalY := len(grid[0])-2, 1
	turnCost := 1000

	cost := cheapestPathCost(&grid, startX, startY, goalX, goalY, turnCost)
	return cost
}

// part two
func part2(input string) int {
	cost, path := alternative(input)
	fmt.Println(cost)

	return path
}

// Heuristic calculates Manhattan distance to the goal
func Heuristic(x, y, goalX, goalY int) int {
	return int(math.Abs(float64(x-goalX)) + math.Abs(float64(y-goalY)))
}

// TurnPenalty calculates the cost of turning from one direction to another
func TurnPenalty(currDir, nextDir string, directions *[]string) int {
	currIdx := indexOf(*directions, currDir)
	nextIdx := indexOf(*directions, nextDir)
	if (nextIdx-currIdx+4)%4 == 2 { // Opposite direction
		return 2
	}
	if currIdx != nextIdx { // 90-degree turn
		return 1
	}
	return 0
}

func indexOf(arr []string, val string) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

// cheapestPathCost finds the cheapest path in the grid considering turn costs
func cheapestPathCost(grid *[][]rune, startX, startY, goalX, goalY, turnCost int) int {
	// Directional constants
	var directions = []string{"N", "E", "S", "W"}
	var dxDy = map[string][2]int{
		"N": {0, -1},
		"E": {1, 0},
		"S": {0, 1},
		"W": {-1, 0},
	}
	pq := &PriorityQueue{}
	heap.Init(pq)

	startState := &State{
		x:         startX,
		y:         startY,
		direction: "E",
		cost:      0,
		h:         Heuristic(startX, startY, goalX, goalY),
	}
	heap.Push(pq, startState)

	visited := make(map[Position]int)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*State)
		key := Position{x: curr.x, y: curr.y, direction: curr.direction}

		// Skip if a better cost was already recorded
		if prevCost, ok := visited[key]; ok && prevCost <= curr.cost {
			continue
		}
		visited[key] = curr.cost

		// Goal check
		if curr.x == goalX && curr.y == goalY {
			return curr.cost
		}

		// Explore neighbors
		for _, nextDir := range directions {
			nx, ny := curr.x+dxDy[nextDir][0], curr.y+dxDy[nextDir][1]

			// Boundary and obstacle check
			if (*grid)[nx][ny] == '#' {
				continue
			}

			moveCost := 1 // Base movement cost
			turnPenalty := TurnPenalty(curr.direction, nextDir, &directions) * turnCost
			nextCost := curr.cost + moveCost + turnPenalty

			nextState := &State{
				x:         nx,
				y:         ny,
				direction: nextDir,
				cost:      nextCost,
				h:         Heuristic(nx, ny, goalX, goalY),
			}

			heap.Push(pq, nextState)
		}
	}

	return -1 // No path found
}

type state struct {
	Pos image.Point
	Dir image.Point
}

type QI struct {
	State state
	Cost  int
	Path  map[image.Point]struct{}
}

func alternative(input string) (int, int) {
	var start image.Point
	grid := map[image.Point]rune{}
	for y, s := range strings.Fields(input) {
		for x, r := range s {
			if r == 'S' {
				start = image.Point{X: x, Y: y}
			}
			grid[image.Point{X: x, Y: y}] = r
		}
	}

	dist := map[state]int{}
	queue := []QI{{
		state{start, image.Point{X: 1}},
		0,
		map[image.Point]struct{}{start: {}},
	}}

	cost, path := math.MaxInt, map[image.Point]struct{}{}
	for len(queue) > 0 {
		slices.SortFunc(queue, func(a, b QI) int {
			return cmp.Compare(a.Cost, b.Cost)
		})
		i := queue[0]
		queue = queue[1:]

		if c, ok := dist[i.State]; ok && c < i.Cost {
			continue
		}
		dist[i.State] = i.Cost

		if grid[i.State.Pos] == 'E' && i.Cost <= cost {
			cost = i.Cost
			maps.Copy(path, i.Path)
		}

		for d, c := range map[image.Point]int{
			i.State.Dir:                     1,
			{-i.State.Dir.Y, i.State.Dir.X}: 1001,
			{i.State.Dir.Y, -i.State.Dir.X}: 1001,
		} {
			n := state{i.State.Pos.Add(d), d}
			if grid[n.Pos] == '#' {
				continue
			}
			p := maps.Clone(i.Path)
			p[n.Pos] = struct{}{}
			queue = append(queue, QI{n, i.Cost + c, p})
		}
	}

	return cost, len(path)
}

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day24.txt")
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
func part1(input string) uint {

	blocks, err := utils.ParseBlocksOfLines(input)
	utils.HandleErr(err)

	wires := map[string]int{}
	for _, line := range blocks[0] {
		parts := strings.Fields(line)
		wires[strings.Trim(parts[0], ":")] = int(parts[1][0] - '0')
	}

	gates := map[string]int{}
	queue := make([][5]string, 0)

	for _, line := range blocks[1] {
		parts := strings.Fields(line)
		if strings.HasPrefix(parts[0], "x") || strings.HasPrefix(parts[0], "y") {
			a, ok := wires[parts[0]]
			b, ok2 := wires[parts[2]]
			if ok && ok2 {
				gates[parts[4]] = gate(a, b, parts[1])
			}
		} else {
			queue = append(queue, [5]string{parts[0], parts[1], parts[2], parts[3], parts[4]})
		}
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		a, ok := gates[item[0]]
		b, ok2 := gates[item[2]]
		if ok && ok2 {
			gates[item[4]] = gate(a, b, item[1])
		} else {
			queue = append(queue, item)
		}
	}

	var res uint = 0
	for k, v := range gates {
		if strings.HasPrefix(k, "z") {
			var idx int
			_, _ = fmt.Sscanf(k, "z%d", &idx)
			res |= uint(v) << idx
		}
	}

	return res
}

// part two
func part2(input string) int {
	return 0
}

func gate(a, b int, op string) int {
	switch op {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "XOR":
		return a ^ b
	default:
		return 0
	}
}

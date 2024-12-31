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
	queue := make([][4]string, 0)

	for _, line := range blocks[1] {
		parts := strings.Fields(line)
		if strings.HasPrefix(parts[0], "x") || strings.HasPrefix(parts[0], "y") {
			a, ok := wires[parts[0]]
			b, ok2 := wires[parts[2]]
			if ok && ok2 {
				gates[parts[4]] = gate(a, b, parts[1])
			}
		} else {
			queue = append(queue, [4]string{parts[0], parts[1], parts[2], parts[4]})
		}
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		a, ok := gates[item[0]]
		b, ok2 := gates[item[2]]
		if ok && ok2 {
			gates[item[3]] = gate(a, b, item[1])
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

// part two
func part2(input string) int {
	blocks, err := utils.ParseBlocksOfLines(input)
	utils.HandleErr(err)

	gates := make(map[[3]string]string)
	for _, line := range blocks[1] {
		parts := strings.Fields(line)
		gates[[3]string{parts[0], parts[2], parts[1]}] = parts[4]
	}

	swaps := map[string]string{}
	results := make([]string, 0)

	repair(gates, 0, swaps, "", &results)

	return 0
}

const EndLevel = 45
const NotFound = "***"

func repair(gates map[[3]string]string, currentLevel int, swaps map[string]string, prevOutput string, results *[]string) bool {
	findSwap := func(v1 string) string {
		if val, exists := swaps[v1]; exists {
			return val
		}
		return v1
	}

	findGate := func(v1, v2, g string) string {
		t1 := [3]string{v1, v2, g}
		t2 := [3]string{v2, v1, g}
		if res, exists := gates[t1]; exists {
			return findSwap(res)
		}
		if res, exists := gates[t2]; exists {
			return findSwap(res)
		}
		return NotFound
	}

	isComparisonValid := func(zbit, xor2, and2, or1 string) bool {
		return zbit == xor2 && zbit != and2 && zbit != or1 && findIndex(*results, NotFound) == -1
	}

	swapRepair := func(or1 string, wires2 []string) bool {
		for i := 0; i < len(wires2); i++ {
			for j := i + 1; j < len(wires2); j++ {
				swapsCopy := make(map[string]string)
				for k, v := range swaps {
					swapsCopy[k] = v
				}
				swapsCopy[wires2[i]] = wires2[j]
				swapsCopy[wires2[j]] = wires2[i]

				var resultsCopy []string
				if repair(gates, currentLevel+1, swapsCopy, or1, &resultsCopy) {
					return true
				}
			}
		}
		return false
	}

	if currentLevel == EndLevel {
		res := make([]string, 0)
		for k := range swaps {
			res = append(res, k)
		}
		fmt.Println(strings.Join(res, ","))
		return true
	}

	if currentLevel == 0 {
		startGate := findGate("x00", "y00", "AND")
		return repair(gates, 1, swaps, startGate, results)
	}

	levelStr := fmt.Sprintf("%02d", currentLevel)
	xbit := "x" + levelStr
	ybit := "y" + levelStr
	zbit := "z" + levelStr
	xor1 := findGate(xbit, ybit, "XOR")
	and1 := findGate(xbit, ybit, "AND")
	xor2 := findGate(xor1, prevOutput, "XOR")
	and2 := findGate(xor1, prevOutput, "AND")
	or1 := findGate(and1, and2, "OR")

	if !isComparisonValid(zbit, xor2, and2, or1) {
		*results = append(*results, xor1, and1, xor2, and2, or1, zbit)
		return false
	}

	var intermediateResults []string
	if !repair(gates, currentLevel+1, swaps, or1, &intermediateResults) {
		return swapRepair(or1, intermediateResults)
	}

	return true
}

func findIndex(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

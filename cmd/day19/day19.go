package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

// main is the entry point of the program. It reads input from a file, processes parts one and two, and prints the results.
func main() {
	// read form file
	input, err := utils.ReadFile("resources/day19.txt")
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

// part1 takes a string input, initializes data, and computes the first result using the solve function logic.
func part1(input string) int {
	towels, designs := initialize(input)

	res, _ := solve(towels, designs)

	return res
}

// part2 processes the input string, calculates the total number of ways to use towels for all designs, and returns the result.
func part2(input string) int {
	towels, designs := initialize(input)

	_, res := solve(towels, designs)

	return res
}

// initialize parses the input string into two slices: towels and designs, trimming whitespace from their elements.
func initialize(input string) ([]string, []string) {
	blocks, err := utils.ParseBlocksOfLines(input)
	utils.HandleErr(err)

	towels := strings.Split(blocks[0][0], ",")
	for i, t := range towels {
		towels[i] = strings.TrimSpace(t)
	}

	designs := blocks[1]
	for i, d := range designs {
		designs[i] = strings.TrimSpace(d)
	}

	return towels, designs
}

// solve computes the number of feasible designs and the total ways to create them.
// It takes two slices of strings: towels and designs, and returns two integers:
//   - The count of designs that can be created using the towels.
//   - The total number of combinations to create all feasible designs.
func solve(towels, designs []string) (int, int) {
	// Declare the recursive function `ways`, which calculates the number of ways
	// to create a specific design.
	var ways func(string) int

	// Create a cache (memoization) map to store already computed results for a design.
	cache := map[string]int{}

	// Define the recursive `ways` function.
	ways = func(design string) (n int) {
		// Check if the result for the current design is already computed and in the cache.
		if n, ok := cache[design]; ok {
			return n
		}

		// Use `defer` to ensure that once the function completes,
		// the result for the current design is stored in the cache.
		defer func() { cache[design] = n }()

		// Base case: If the design is an empty string, it can only be formed in one way (by doing nothing).
		if design == "" {
			return 1
		}

		// Iterate over each towel string.
		for _, s := range towels {
			// Check if the `design` starts with the current towel string `s`.
			if strings.HasPrefix(design, s) {
				// If so, recursively calculate the number of ways for the remaining part of the design
				// (after removing the prefix `s`) and add it to `n`.
				n += ways(design[len(s):])
			}
		}

		// Return the total number of ways to construct the current design.
		return n
	}

	// Initialize two results:
	// `res1` — the count of feasible designs (those with at least one way to construct them).
	// `res2` — the total number of combinations to construct all feasible designs.
	res1, res2 := 0, 0

	// Iterate over each design from the input `designs` slice.
	for _, s := range designs {
		// Compute the number of ways to construct the current design using the `ways` function.
		if w := ways(s); w > 0 {
			// If there is at least one way (`w > 0`), increment the count of feasible designs (`res1`).
			res1++
			// Add the number of ways (`w`) to the total combinations (`res2`).
			res2 += w
		}
	}

	// Return the count of feasible designs and the total number of combinations.
	return res1, res2
}

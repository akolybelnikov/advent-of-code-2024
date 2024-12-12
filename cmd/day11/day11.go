package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

func main() {
	// Read from file
	input, err := utils.ReadFile("resources/day11.txt")
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
	stones := map[int]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), " ") {
		n, err := strconv.Atoi(s)
		utils.HandleErr(err)
		stones[n]++
	}

	return run(stones, 25)
}

// part two
func part2(input string) int {
	stones := map[int]int{}
	for _, s := range strings.Split(strings.TrimSpace(input), " ") {
		n, err := strconv.Atoi(s)
		utils.HandleErr(err)
		stones[n]++
	}

	return run(stones, 75)
}

func run(stones map[int]int, blinks int) (r int) {
	for range blinks {
		next := map[int]int{}
		for k, v := range stones {
			if k == 0 {
				next[1] += v
			} else if s := strconv.Itoa(k); len(s)%2 == 0 {
				n1, _ := strconv.Atoi(s[:len(s)/2])
				n2, _ := strconv.Atoi(s[len(s)/2:])
				next[n1] += v
				next[n2] += v
			} else {
				next[k*2024] += v
			}
		}
		stones = next
	}
	for _, v := range stones {
		r += v
	}
	return r
}

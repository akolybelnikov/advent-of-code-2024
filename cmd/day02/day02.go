package main

import (
	"fmt"
	"math"
	"os"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day02.txt")
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

	intSlices, err := parseLines(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, slice := range intSlices {
		sum += inspectIntSlice(slice, false)
	}

	return sum
}

// part two
func part2(input string) int {
	var sum int

	intSlices, err := parseLines(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, slice := range intSlices {
		sum += inspectIntSlice(slice, true)
	}

	return sum
}

func parseLines(input string) ([][]int, error) {
	lines, err := utils.ParseLines(input)
	if err != nil {
		return nil, err
	}

	intSlices, err := utils.ConvertLinesToIntSlices(lines)
	if err != nil {
		return nil, err
	}

	return intSlices, nil
}

func inspectIntSlice(intSlice []int, canNormalize bool) int {
	if isConsistent(intSlice) {
		if isValid(intSlice) {
			return 1
		}
	}

	if canNormalize {
		for i := range intSlice {
			newSlice := make([]int, len(intSlice))
			copy(newSlice, intSlice)
			testSlice := append(newSlice[:i], newSlice[i+1:]...)
			if isConsistent(testSlice) {
				if isValid(testSlice) {
					return 1
				}
			}
		}
	}

	return 0
}

func isConsistent(intSlice []int) bool {
	countViolations := 0

	for i := 1; i < len(intSlice); i++ {
		if intSlice[i] <= intSlice[i-1] {
			countViolations++
			break
		}
	}

	if countViolations == 0 {
		return true
	}

	countViolations = 0

	for i := 1; i < len(intSlice); i++ {
		if intSlice[i] >= intSlice[i-1] {
			countViolations++
			break
		}
	}

	if countViolations == 0 {
		return true
	}

	return false
}

func isValid(intSlice []int) bool {
	countViolations := 0
	for i := 1; i < len(intSlice); i++ {
		if math.Abs(float64(intSlice[i]-intSlice[i-1])) > 3 {
			countViolations++
			break
		}
	}
	return countViolations == 0
}

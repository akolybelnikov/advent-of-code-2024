package main

import (
	"fmt"
	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day03.txt")
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

func findSumOfProducts(input string) (int, error) {
	var sum int
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		num1, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}
		num2, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}
		sum += num1 * num2
	}

	return sum, nil
}

// part one
func part1(input string) int {
	sum, err := findSumOfProducts(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return sum
}

// part two
func part2(input string) int {
	var sum int
	splitByDont := strings.Split(input, "don't()")
	validParts := make([]string, 0)
	for i, split := range splitByDont {
		if i == 0 {
			validParts = append(validParts, split)
		}
		splitByDo := strings.Split(split, "do()")
		if len(splitByDo) > 1 {
			validParts = append(validParts, splitByDo[1:]...)
		}
	}
	for _, part := range validParts {
		intSum, err := findSumOfProducts(part)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum += intSum
	}

	return sum
}

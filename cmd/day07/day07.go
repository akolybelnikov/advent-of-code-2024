package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day07.txt")
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

func createNumbers(line []string) ([]*big.Int, error) {
	var ints []*big.Int
	for _, part := range line {
		n, ok := new(big.Int).SetString(part, 10)
		if !ok {
			return nil, fmt.Errorf(fmt.Sprintf("Cannot convert %s to big.Int", part))
		}
		ints = append(ints, n)
	}

	return ints, nil
}

func parseLine(line string) (*big.Int, []*big.Int, error) {
	parts := strings.Split(line, ":")
	target, ok := new(big.Int).SetString(parts[0], 10)
	if !ok {
		return nil, nil, fmt.Errorf("error parsing number")
	}

	numbers, err := createNumbers(strings.Fields(parts[1]))
	if err != nil {
		return nil, nil, err
	}

	return target, numbers, nil
}

func hasSolutions(arr *[]*big.Int, target *big.Int, index int, cur *big.Int) bool {
	if index == len(*arr) {
		return cur.Cmp(target) == 0
	}

	if cur.Cmp(target) == 1 {
		return false
	}

	if hasSolutions(arr, target, index+1, new(big.Int).Add(cur, (*arr)[index])) {
		return true
	}

	if hasSolutions(arr, target, index+1, new(big.Int).Mul(cur, (*arr)[index])) {
		return true
	}

	return false
}

// part one
func part1(input string) int64 {
	lines, err := utils.ParseLines(input)
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}

	sum := new(big.Int)

	for _, line := range lines {
		target, numbers, err2 := parseLine(line)
		if err2 != nil {
			fmt.Println("Error parsing line: ", err2)
			os.Exit(1)
		}

		if hasSolutions(&numbers, target, 1, numbers[0]) {
			sum = new(big.Int).Add(sum, target)
		}
	}

	return sum.Int64()
}

func hasSolutions2(arr *[]*big.Int, target *big.Int, index int, cur *big.Int) bool {
	if index == len(*arr) {
		return cur.Cmp(target) == 0
	}

	if cur.Cmp(target) == 1 {
		return false
	}

	if hasSolutions2(arr, target, index+1, new(big.Int).Add(cur, (*arr)[index])) {
		return true
	}

	if hasSolutions2(arr, target, index+1, new(big.Int).Mul(cur, (*arr)[index])) {
		return true
	}

	concatenated := cur.String() + (*arr)[index].String()
	val, ok := new(big.Int).SetString(concatenated, 10)
	if !ok {
		fmt.Println("Error concatenating numbers: ", concatenated)
		os.Exit(1)
	}
	if hasSolutions2(arr, target, index+1, val) {
		return true
	}

	return false
}

// part two
func part2(input string) int64 {
	lines, err := utils.ParseLines(input)
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}

	sum := new(big.Int)

	for _, line := range lines {
		target, numbers, err2 := parseLine(line)
		if err2 != nil {
			fmt.Println("Error parsing line: ", err2)
			os.Exit(1)
		}

		if hasSolutions2(&numbers, target, 1, numbers[0]) {
			sum = new(big.Int).Add(sum, target)
		}
	}

	return sum.Int64()
}

package main

import (
	"fmt"
	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
	"math"
	"os"
	"sort"
)

func main() {
	input, err := utils.ReadFile("resources/day01.txt")
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
	lines, err := utils.ParseLines(input)
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}

	arr1 := make([]int, len(lines))
	arr2 := make([]int, len(lines))

	intSlices, err := utils.ConvertLinesToIntSlices(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, slice := range intSlices {
		arr1[i] = slice[0]
		arr2[i] = slice[1]
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i, i2 := range arr1 {
		sum += int(math.Abs(float64(arr2[i] - i2)))
	}

	return sum
}

// part two
func part2(input string) int {
	var sum int
	lines, err := utils.ParseLines(input)
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}

	arr1 := make([]int, len(lines))
	map2 := make(map[int]int)

	intSlices, err := utils.ConvertLinesToIntSlices(lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, slice := range intSlices {
		arr1[i] = slice[0]
		map2[slice[1]]++
	}

	for _, i2 := range arr1 {
		num, ok := map2[i2]
		if ok {
			sum += i2 * num
		}
	}

	return sum
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
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
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}
	arr1 := make([]int, len(lines))
	arr2 := make([]int, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		nums := strings.Fields(line)
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		arr1[i] = num1
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		arr2[i] = num2
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
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}
	arr1 := make([]int, len(lines))
	map2 := make(map[int]int)

	for i, line := range lines {
		if line == "" {
			continue
		}
		nums := strings.Fields(line)
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		arr1[i] = num1
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		map2[num2]++
	}

	for _, i2 := range arr1 {
		num, ok := map2[i2]
		if ok {
			sum += i2 * num
		}
	}

	return sum
}

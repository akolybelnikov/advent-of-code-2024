package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

func main() {
	input, err := utils.ReadFile("resources/day05.txt")
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

func makeRules(input []string) map[string][]string {
	var rules = make(map[string][]string)
	for _, rule := range input {
		split := strings.Split(rule, "|")
		rules[split[0]] = append(rules[split[0]], split[1])
	}

	return rules
}

// part one
func part1(input string) int {

	blocks, err := utils.ParseBlocksOfLines(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rules := makeRules(blocks[0])

	var sum int
	for _, report := range blocks[1] {
		sum += inspectReport(report, &rules)
	}

	return sum
}

func followsRules(nums []string, rules *map[string][]string) bool {
	for i, num := range nums {
		next, ok := (*rules)[num]
		if !ok && i != len(nums)-1 {
			return false
		}
		for j := i + 1; j < len(nums); j++ {
			found := false
			for _, nextNum := range next {
				if nextNum == nums[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}

	return true
}

func followsRulesReverse(nums []string, rules *map[string][]string) bool {
	for i := len(nums) - 1; i >= 0; i-- {
		prev, ok := (*rules)[nums[i]]
		if !ok && i != len(nums)-1 {
			return false
		}
		for j := i - 1; j >= 0; j-- {
			found := false
			for _, prevNum := range prev {
				if prevNum == nums[j] {
					found = true
					break
				}
			}
			if found {
				return false
			}
		}
	}

	return true
}

func inspectReport(report string, rules *map[string][]string) int {
	nums := strings.Split(report, ",")
	if !followsRules(nums, rules) || !followsRulesReverse(nums, rules) {
		return 0
	}

	midNumStr := len(nums) / 2
	midNum, err := strconv.Atoi(nums[midNumStr])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return midNum
}

// part two
func part2(input string) int {
	blocks, err := utils.ParseBlocksOfLines(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rules := makeRules(blocks[0])

	var sum int
	for _, report := range blocks[1] {
		if inspectReport(report, &rules) == 0 {
			nums := strings.Split(report, ",")
			ranks := make(map[string]int)
			for i, num := range nums {
				ranks[num] = 0
				for j := 0; j < len(nums); j++ {
					if j != i {
						next, ok := rules[nums[j]]
						if !ok {
							continue
						}
						for _, nextNum := range next {
							if nextNum == num {
								ranks[num]++
								break
							}
						}
					}
				}
			}
			for numStr, rank := range ranks {
				if rank == len(nums)/2 {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					sum += num
				}
			}
		}
	}

	return sum
}

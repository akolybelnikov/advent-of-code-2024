package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day22.txt")
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
	var ans uint
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	for _, line := range lines {
		num, numErr := strconv.ParseUint(line, 10, 64)
		utils.HandleErr(numErr)

		ans += evolve2000(uint(num))
	}

	return ans
}

// part two
func part2(input string) int {
	return 0
}

func evolve(secret uint) uint {
	// Shift left by 6 and XOR
	mul := secret << 6
	secret = secret ^ mul

	// Mask lower 24 bits
	secret = secret & 0xFFFFFF

	// Shift right by 15 and XOR
	div := secret >> 5
	secret = secret ^ div

	// Mask lower 24 bits again
	secret = secret & 0xFFFFFF

	// Shift left by 11 and XOR
	mul = secret << 11
	secret = secret ^ mul

	// Final mask to lower 24 bits
	secret = secret & 0xFFFFFF

	return secret
}

func evolve2000(secret uint) uint {
	for i := 0; i < 2000; i++ {
		secret = evolve(secret)
	}

	return secret
}

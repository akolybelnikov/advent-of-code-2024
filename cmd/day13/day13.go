package main

import (
	"fmt"
	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
	"image"
	"os"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day13.txt")
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
	blocks, err := utils.ParseBlocksOfLines(input)
	utils.HandleErr(err)

	ans1, _ := solve(blocks)

	return ans1
}

// part two
func part2(input string) int {
	blocks, err := utils.ParseBlocksOfLines(input)
	utils.HandleErr(err)

	_, ans2 := solve(blocks)

	return ans2
}

func solve(blocks [][]string) (int, int) {
	ans1, ans2 := 0, 0
	for _, block := range blocks {
		var a, b, c image.Point
		_, _ = fmt.Sscanf(block[0], "Button A: X+%d, Y+%d", &a.X, &a.Y)
		_, _ = fmt.Sscanf(block[1], "Button B: X+%d, Y+%d", &b.X, &b.Y, &c.X, &c.Y)
		_, _ = fmt.Sscanf(block[2], "Prize: X=%d, Y=%d", &c.X, &c.Y)
		ans1 += calc(a, b, c)
		ans2 += calc(a, b, c.Add(image.Point{X: 10000000000000, Y: 10000000000000}))
	}

	return ans1, ans2
}

func calc(a, b, c image.Point) int {
	var ap int
	diffA := a.X*b.Y - a.Y*b.X
	if diffA != 0 {
		ap = (b.Y*c.X - b.X*c.Y) / diffA
	}

	var bp int
	diffB := a.Y*b.X - a.X*b.Y
	if diffB != 0 {
		bp = (a.Y*c.X - a.X*c.Y) / diffB
	}

	check := a.Mul(ap).Add(b.Mul(bp))
	if check == c {
		return ap*3 + bp
	}
	return 0
}

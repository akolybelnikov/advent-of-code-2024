package main

import (
	"fmt"
	"os"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

type pin [5]int8
type combination struct {
	lock    pin
	key     pin
	overlap bool
}

func (c *combination) fit() {
	for i := 0; i < 5; i++ {
		if c.lock[i]+c.key[i] > 5 {
			c.overlap = true
			break
		}
	}
}

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day25.txt")
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
	res := 0
	blocks, err := utils.ParseBlocksOfLines(input)
	utils.HandleErr(err)

	locks, keys := make([]pin, 0), make([]pin, 0)
	for _, block := range blocks {
		if block[0][0] == '#' {
			locks = append(locks, lockToPin(block))
		} else {
			keys = append(keys, keyToPin(block))
		}
	}

	for _, lock := range locks {
		for _, key := range keys {
			cb := combination{lock, key, false}
			cb.fit()
			if !cb.overlap {
				res++
			}
		}
	}

	return res
}

// part two
func part2(input string) int {
	return 0
}

func lockToPin(lock []string) pin {
	p := pin{}
	for _, s := range lock[1 : len(lock)-1] {
		r := []rune(s)
		for j, v := range r {
			if v == '#' {
				p[j]++
			}
		}
	}
	return p
}

func keyToPin(key []string) pin {
	p := pin{}
	for i := len(key) - 2; i > 0; i-- {
		r := []rune(key[i])
		for j, v := range r {
			if v == '#' {
				p[j]++
			}
		}
	}
	return p
}

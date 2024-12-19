package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

type Register struct {
	A, B, C int
	program []int
	output  []string
}

func (r *Register) combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return r.A
	case 5:
		return r.B
	case 6:
		return r.C
	default:
		fmt.Println("Invalid operand", operand)
		os.Exit(1)
		return -1
	}
}

func (r *Register) instruct(instruction int, operand int, idx int) int {
	switch instruction {
	case 0:
		r.A = r.A / (1 << r.combo(operand))
	case 1:
		r.B = r.B ^ operand
	case 2:
		r.B = r.combo(operand) % 8
	case 3:
		if r.A > 0 {
			return operand
		}
	case 4:
		r.B = r.B ^ r.C
	case 5:
		r.output = append(r.output, fmt.Sprintf("%d", r.combo(operand)%8))
	case 6:
		r.B = r.A / (1 << r.combo(operand))
	case 7:
		r.C = r.A / (1 << r.combo(operand))
	default:
		fmt.Println("Invalid instruction", instruction)
		os.Exit(1)
		return -1
	}

	return idx + 2
}

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day17.txt")
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
func part1(input string) string {
	parsed, err := utils.ParseLines(input)
	utils.HandleErr(err)

	r := initialize(parsed)

	idx := 0
	for idx < len(r.program)-1 {
		instruction, operand := r.program[idx], r.program[idx+1]
		idx = r.instruct(instruction, operand, idx)
	}

	return strings.Join(r.output, ",")
}

// part two
func part2(input string) int {
	// Parse the program and initialize the Register
	parsed, err := utils.ParseLines(input)
	utils.HandleErr(err)
	r := initialize(parsed)

	a := 0
	for n := len(r.program) - 1; n >= 0; n-- {
		a <<= 3
		for !slices.Equal(run(a, r.B, r.C, r.program), r.program[n:]) {
			a++
		}
	}
	fmt.Println(a)

	return a
}

func initialize(input []string) Register {
	register := Register{}
	_, _ = fmt.Sscanf(input[0], "Register A: %d", &register.A)
	_, _ = fmt.Sscanf(input[1], "Register B: %d", &register.B)
	_, _ = fmt.Sscanf(input[2], "Register C: %d", &register.C)

	register.program = make([]int, 0)
	programStr := strings.Fields(input[3])[1]
	for _, numStr := range strings.Split(programStr, ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			utils.HandleErr(err)
		}
		register.program = append(register.program, num)
	}

	return register
}

func run(a, b, c int, pgm []int) (out []int) {
	for ip := 0; ip < len(pgm); ip += 2 {
		literal := pgm[ip+1]
		combo := []int{0, 1, 2, 3, a, b, c, 7}[literal]

		switch pgm[ip] {
		case 0:
			a >>= combo
		case 1:
			b ^= literal
		case 2:
			b = combo % 8
		case 3:
			if a != 0 {
				ip = literal - 2
			}
		case 4:
			b ^= c
		case 5:
			out = append(out, combo%8)
		case 6:
			b = a >> combo
		case 7:
			c = a >> combo
		}
	}
	return out
}

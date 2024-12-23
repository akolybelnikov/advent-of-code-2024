package main

import (
	"fmt"
	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
	"os"
	"strconv"
	"strings"
)

type area struct {
	keypad     map[rune]map[rune][]rune
	directions map[rune]map[rune][]rune
}

func (a *area) sequence(curKey rune, nextKey rune) string {
	builder := strings.Builder{}
	numCmd := a.keypad[curKey][nextKey]
	//fmt.Println("numCmd: ", string(numCmd))
	curDir, curDirDir := 'A', 'A'

	for _, n := range numCmd {
		dirCmd := a.directions[curDir][n]
		//fmt.Println("n ", string(n), " | ", "dirCmd: ", string(dirCmd))
		for _, d := range dirCmd {
			dirDirCmd := a.directions[curDirDir][d]
			//fmt.Println("d ", string(d), " | ", "dirDirCmd: ", string(dirDirCmd))
			for _, dir := range dirDirCmd {
				builder.WriteRune(dir)
			}
			curDirDir = d
		}
		curDir = n
	}

	return builder.String()
}

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day21.txt")
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
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	sequences := make(map[string]string)
	a := newArea()
	b := strings.Builder{}

	for _, line := range lines {
		runes := []rune(line)
		runes = append([]rune{'A'}, runes...)
		for i := 0; i < len(runes)-1; i++ {
			b.WriteString(a.sequence(runes[i], runes[i+1]))
		}
		sequences[line] = b.String()
		b.Reset()
	}

	return calculate(sequences)
}

// part two
func part2(input string) int {
	return 0
}

func calculate(sequences map[string]string) int {
	ans := 0
	for code, seq := range sequences {
		num, err := strconv.Atoi(code[:len(code)-1])
		utils.HandleErr(err)
		ans += num * len(seq)
		//fmt.Println(code, num, len(seq), num*len(seq))
	}

	return ans
}

func newArea() *area {
	return &area{
		keypad:     keypad(),
		directions: directions(),
	}
}

func keypad() map[rune]map[rune][]rune {
	return map[rune]map[rune][]rune{
		'A': {
			'0': {'<', 'A'},
			'1': {'^', '<', '<', 'A'},
			'2': {'^', '<', 'A'},
			'3': {'^', 'A'},
			'4': {'^', '^', '<', '<', 'A'},
			'5': {'^', '^', '<', 'A'},
			'6': {'^', '^', 'A'},
			'7': {'^', '^', '^', '<', '<', 'A'},
			'8': {'^', '^', '^', '<', 'A'},
			'9': {'^', '^', '^', 'A'},
			'A': {'A'},
		},
		'0': {
			'A': {'>', 'A'},
			'1': {'^', '<', 'A'},
			'2': {'^', 'A'},
			'3': {'^', '>', 'A'},
			'4': {'^', '^', '<', 'A'},
			'5': {'^', '^', 'A'},
			'6': {'^', '^', '>', 'A'},
			'7': {'^', '^', '^', '<', 'A'},
			'8': {'^', '^', '^', 'A'},
			'9': {'^', '^', '^', '>', 'A'},
			'0': {'A'},
		},
		'1': {
			'A': {'>', 'v', '>', 'A'},
			'0': {'>', 'v', 'A'},
			'2': {'>', 'A'},
			'3': {'>', '>', 'A'},
			'4': {'^', 'A'},
			'5': {'^', '>', 'A'},
			'6': {'^', '>', '>', 'A'},
			'7': {'^', '^', 'A'},
			'8': {'^', '^', '>', 'A'},
			'9': {'^', '^', '>', '>', 'A'},
			'1': {'A'},
		},
		'2': {
			'A': {'v', '>', 'A'},
			'0': {'v', 'A'},
			'1': {'<', 'A'},
			'3': {'>', 'A'},
			'4': {'^', '<', 'A'},
			'5': {'^', 'A'},
			'6': {'^', '>', 'A'},
			'7': {'^', '^', '<', 'A'},
			'8': {'^', '^', 'A'},
			'9': {'>', '^', '^', 'A'},
			'2': {'A'},
		},
		'3': {
			'A': {'v', 'A'},
			'0': {'v', '<', 'A'},
			'1': {'<', '<', 'A'},
			'2': {'<', 'A'},
			'4': {'^', '<', '<', 'A'},
			'5': {'^', '<', 'A'},
			'6': {'^', 'A'},
			'7': {'<', '<', '^', '^', 'A'},
			'8': {'^', '^', '<', 'A'},
			'9': {'^', '^', 'A'},
			'3': {'A'},
		},
		'4': {
			'A': {'v', '>', '>', 'v', 'A'},
			'0': {'v', '>', 'v', 'A'},
			'1': {'v', 'A'},
			'2': {'v', '>', 'A'},
			'3': {'v', '>', '>', 'A'},
			'5': {'>', 'A'},
			'6': {'>', '>', 'A'},
			'7': {'^', 'A'},
			'8': {'^', '>', 'A'},
			'9': {'^', '>', '>', 'A'},
			'4': {'A'},
		},
		'5': {
			'A': {'v', '>', 'v', 'A'},
			'0': {'v', 'v', 'A'},
			'1': {'v', '<', 'A'},
			'2': {'v', 'A'},
			'3': {'v', '>', 'A'},
			'4': {'<', 'A'},
			'6': {'>', 'A'},
			'7': {'^', '<', 'A'},
			'8': {'^', 'A'},
			'9': {'^', '>', 'A'},
			'5': {'A'},
		},
		'6': {
			'A': {'v', 'v', 'A'},
			'0': {'v', 'v', '<', 'A'},
			'1': {'v', '<', '<', 'A'},
			'2': {'v', '<', 'A'},
			'3': {'v', 'A'},
			'4': {'<', '<', 'A'},
			'5': {'<', 'A'},
			'7': {'^', '<', '<', 'A'},
			'8': {'^', '<', 'A'},
			'9': {'^', 'A'},
			'6': {'A'},
		},
		'7': {
			'A': {'v', 'v', '>', '>', 'v', 'A'},
			'0': {'v', 'v', '>', 'v', 'A'},
			'1': {'v', 'v', 'A'},
			'2': {'v', 'v', '>', 'A'},
			'3': {'v', 'v', '>', '>', 'A'},
			'4': {'v', 'A'},
			'5': {'v', '>', 'A'},
			'6': {'v', '>', '>', 'A'},
			'8': {'>', 'A'},
			'9': {'>', '>', 'A'},
			'7': {'A'},
		},
		'8': {
			'A': {'v', 'v', 'v', '>', 'A'},
			'0': {'v', 'v', 'v', 'A'},
			'1': {'v', 'v', '<', 'A'},
			'2': {'v', 'v', 'A'},
			'3': {'v', 'v', '>', 'A'},
			'4': {'v', '<', 'A'},
			'5': {'v', 'A'},
			'6': {'v', '>', 'A'},
			'7': {'<', 'A'},
			'9': {'>', 'A'},
			'8': {'A'},
		},
		'9': {
			'A': {'v', 'v', 'v', 'A'},
			'0': {'v', 'v', 'v', '<', 'A'},
			'1': {'v', 'v', '<', '<', 'A'},
			'2': {'v', 'v', '<', 'A'},
			'3': {'v', 'v', 'A'},
			'4': {'v', '<', '<', 'A'},
			'5': {'v', '<', 'A'},
			'6': {'v', 'A'},
			'7': {'<', '<', 'A'},
			'8': {'<', 'A'},
			'9': {'A'},
		},
	}
}

func directions() map[rune]map[rune][]rune {
	return map[rune]map[rune][]rune{
		'A': {
			'^': {'<', 'A'},
			'v': {'<', 'v', 'A'},
			'<': {'v', '<', '<', 'A'},
			'>': {'v', 'A'},
			'A': {'A'},
		},
		'^': {
			'A': {'>', 'A'},
			'v': {'v', 'A'},
			'<': {'v', '<', 'A'},
			'>': {'v', '>', 'A'},
			'^': {'A'},
		},
		'v': {
			'A': {'>', '^', 'A'},
			'^': {'^', 'A'},
			'<': {'<', 'A'},
			'>': {'>', 'A'},
			'v': {'A'},
		},
		'<': {
			'A': {'>', '>', '^', 'A'},
			'^': {'>', '^', 'A'},
			'v': {'>', 'A'},
			'>': {'>', '>', 'A'},
			'<': {'A'},
		},
		'>': {
			'A': {'^', 'A'},
			'^': {'<', '^', 'A'},
			'<': {'<', '<', 'A'},
			'v': {'<', 'A'},
			'>': {'A'},
		},
	}
}

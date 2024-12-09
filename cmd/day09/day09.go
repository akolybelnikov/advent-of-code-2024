package main

import (
	"fmt"
	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
	"os"
	"slices"
	"strings"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day09.txt")
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
	disc, empty := inputToDisc(input)
	count := 0
	for i := len(disc) - 1; i > 0; i-- {
		if count > len(empty)-1 {
			break
		}
		if disc[i] != -1 {
			disc[empty[count]] = disc[i]
			count++
		}
	}
	disc = disc[:len(disc)-len(empty)]

	return calculateChecksum(&disc)
}

type file struct {
	id   int
	size int
}

// part two
func part2(input string) int {
	diskmap := strings.TrimSpace(string(input)) + "0"

	var fs2 []file
	for id := 0; id*2 < len(diskmap); id++ {
		size, free := int(diskmap[id*2]-'0'), int(diskmap[id*2+1]-'0')
		fs2 = append(fs2, file{id, size}, file{-1, free})
	}

	return run(fs2)
}

func run(fs []file) (checksum int) {
	for f := len(fs) - 1; f >= 0; f-- {
		for free := 0; free < f; free++ {
			if fs[f].id != -1 && fs[free].id == -1 && fs[free].size >= fs[f].size {
				fs = slices.Insert(fs, free, fs[f])
				fs[f+1].id = -1
				fs[free+1].size = fs[free+1].size - fs[f+1].size
			}
		}
	}
	i := 0
	for _, f := range fs {
		for range f.size {
			if f.id != -1 {
				checksum += i * f.id
			}
			i++
		}
	}
	return checksum
}

func inputToDisc(input string) ([]int, []int) {
	disc := make([]int, 0)
	empty := make([]int, 0)
	index := 0
	for i, s := range input {
		n := int(s - '0')
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				disc = append(disc, int(i)/2)
				index++
			}
		} else {
			for j := 0; j < n; j++ {
				disc = append(disc, -1)
				empty = append(empty, index)
				index++
			}
		}
	}
	return disc, empty
}

func calculateChecksum(disc *[]int) int {
	sum := 0
	for i, id := range *disc {
		sum += int(i) * id
	}

	return sum
}

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

type network map[string]subnetwork
type subnetwork []string

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day23.txt")
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
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	n := initialize(lines)
	s := n.findSubnetworks()
	for _, v := range s {
		res += v
	}

	return res
}

// part two
func part2(input string) string {
	res := ""
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)

	n := initialize(lines)
	s := n.findLargeSubnetworks()
	for k, _ := range s {
		if len(k) > len(res) {
			res = k
		}
	}

	return res
}

func initialize(lines []string) network {
	n := make(map[string]subnetwork)

	for _, line := range lines {
		split := strings.Split(line, "-")
		c1, c2 := split[0], split[1]
		n[c1] = append(n[c1], c2)
		n[c2] = append(n[c2], c1)
	}

	return n
}

func (s *subnetwork) contains(c string) bool {
	for _, i := range *s {
		if i == c {
			return true
		}
	}
	return false
}

func (n *network) findSubnetworks() map[string]int {
	res := make(map[string]int)
	for c, s := range *n {
		for i := 0; i < len(s); i++ {
			s2 := (*n)[s[i]]
			for j := i + 1; j < len(s); j++ {
				if s2.contains(s[j]) {
					sub := []string{c, s[i], s[j]}
					sort.Strings(sub)
					key := strings.Join(sub, ",")
					if strings.HasPrefix(c, "t") || strings.HasPrefix(s[i], "t") || strings.HasPrefix(s[j], "t") {
						res[key] = 1
					} else {
						res[key] = 0
					}

				}
			}
		}
	}

	return res
}

func (n *network) findLargeSubnetworks() map[string]struct{} {
	res := make(map[string]struct{})
	for c, s := range *n {
		for x, i := range s {
			sub := subnetwork{c, i}
			for _, j := range s[x+1:] {
				in := true
				arr := (*n)[j]
				for _, a := range sub {
					if !arr.contains(a) {
						in = false
						break
					}
				}
				if in {
					sub = append(sub, j)
				}
			}
			sort.Strings(sub)
			key := strings.Join(sub, ",")
			res[key] = struct{}{}
		}
	}

	return res
}

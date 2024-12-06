package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func ParseLines(data string) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ParseBlocksOfLines(data string) ([][]string, error) {
	var blocks = make([][]string, 0)
	var curBlock []string
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			curBlock = append(curBlock, line)
		} else {
			if len(curBlock) > 0 {
				blocks = append(blocks, curBlock)
				curBlock = make([]string, 0)
			}
		}
	}

	if len(curBlock) > 0 {
		blocks = append(blocks, curBlock)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return blocks, nil
}

func ConvertLinesToIntSlices(data []string) ([][]int, error) {
	intSlices := make([][]int, len(data))
	for i, line := range data {
		numStrings := strings.Fields(line)
		for _, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			if err != nil {
				return intSlices, err
			}
			intSlices[i] = append(intSlices[i], num)
		}
	}

	return intSlices, nil
}

func ConvertLinesToRuneSlices(data []string) [][]rune {
	runeSlices := make([][]rune, len(data))
	for i, line := range data {
		runs := []rune(line)
		runeSlices[i] = runs
	}
	return runeSlices
}

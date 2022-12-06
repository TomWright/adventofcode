package main

import (
	"fmt"
	"github.com/TomWright/adventofcode/2022/util"
	"io"
	"os"
)

const input = "day06/input.txt"

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	contents, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	input := string(contents)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(input), part2(input))
}

func findStart(input string, requiredChars int) int {
	cur := make([]rune, 0)
outer:
	for k, v := range input {
		if len(cur) < requiredChars {
			cur = append(cur, v)
			continue
		}
		cur = append(cur[1:], v)
		for _, c := range cur {
			if util.ContainsCount(c, cur) > 1 {
				continue outer
			}
		}
		return k
	}
	return -1
}

func part1(input string) int {
	return findStart(input, 4) + 1
}

func part2(input string) int {
	return findStart(input, 14) + 1
}

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

func isDistinct(in []rune) bool {
	for _, r := range in {
		if util.ContainsCount(r, in) > 1 {
			return false
		}
	}
	return true
}

func findStart(input string, requiredChars int) int {
	in := []rune(input)
	cur := in[0:requiredChars]
	in = in[requiredChars:]
	for k, v := range input {
		cur = append(cur[1:], v)
		if !isDistinct(cur) {
			continue
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

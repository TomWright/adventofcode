package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "day04/input.txt"

func NewElf(input string) Elf {
	parts := strings.Split(input, "-")
	from, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	to, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return Elf{
		Min: from,
		Max: to,
	}
}

type Elf struct {
	Min int
	Max int
}

func (e Elf) Contains(other Elf) bool {
	return e.Min <= other.Min && e.Max >= other.Max
}

func (e Elf) Overlaps(other Elf) bool {
	return (other.Min >= e.Min && other.Min <= e.Max) || (other.Max >= e.Min && other.Max <= e.Max)
}

func (e Elf) Sections() []int {
	res := make([]int, 0)
	for i := e.Min; i <= e.Max; i++ {
		res = append(res, i)
	}
	return res
}

func NewPair(input string) Pair {
	parts := strings.Split(input, ",")
	return Pair{
		First:  NewElf(parts[0]),
		Second: NewElf(parts[1]),
	}
}

type Pair struct {
	First  Elf
	Second Elf
}

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pairs := make([]Pair, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()
		pair := NewPair(t)
		pairs = append(pairs, pair)
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(pairs), part2(pairs))
}

func part1(pairs []Pair) int {
	res := 0
	for _, p := range pairs {
		if p.First.Contains(p.Second) || p.Second.Contains(p.First) {
			res++
		}
	}
	return res
}

func part2(pairs []Pair) int {
	res := 0
	for _, p := range pairs {
		if p.First.Overlaps(p.Second) || p.Second.Overlaps(p.First) {
			res++
		}
	}
	return res
}

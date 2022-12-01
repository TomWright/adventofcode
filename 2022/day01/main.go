package main

import (
	"bufio"
	"fmt"
	"github.com/TomWright/adventofcode/2022/util"
	"os"
	"sort"
	"strconv"
)

type Item struct {
	Calories int
}

type Elf struct {
	Items []Item
}

func (e *Elf) TotalCalories() int {
	return util.Sum(
		util.Map(e.Items, func(i Item) int {
			return i.Calories
		}),
	)
}

const input = "input.txt"

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	currentElf := &Elf{Items: nil}

	elfs := make([]*Elf, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			elfs = append(elfs, currentElf)
			currentElf = &Elf{Items: nil}
			continue
		}

		calories, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		currentElf.Items = append(currentElf.Items, Item{Calories: calories})
	}

	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].TotalCalories() > elfs[j].TotalCalories()
	})

	part1 := elfs[0].TotalCalories()

	part2 := util.Sum(
		util.Map(util.First(3, elfs), func(i *Elf) int {
			return i.TotalCalories()
		}),
	)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

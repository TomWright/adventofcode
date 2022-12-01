package main

import (
	"bufio"
	"fmt"
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
	sum := 0
	for _, i := range e.Items {
		sum += i.Calories
	}
	return sum
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

	sum := 0
	for x := 0; x < 3; x++ {
		sum += elfs[x].TotalCalories()
	}

	fmt.Println(sum)
}

package main

import (
	"bufio"
	"fmt"
	"github.com/TomWright/adventofcode/2022/util"
	"os"
)

func NewRucksack(contents string) *Rucksack {
	items := make([]Item, 0)
	for _, i := range contents {
		items = append(items, Item(i))
	}

	numItems := len(items)
	compartmentSize := numItems / 2

	r := &Rucksack{
		Compartments: []Compartment{
			items[:compartmentSize],
			items[compartmentSize:],
		},
	}

	return r
}

type Group []*Rucksack

func (g Group) FindBadge() Item {
	for _, item := range g[0].Items() {
		if g[1].Contains(item) && g[2].Contains(item) {
			return item
		}
	}
	return ""
}

type Rucksack struct {
	Compartments []Compartment
}

func (r *Rucksack) Items() []Item {
	return append(r.Compartments[0], r.Compartments[1]...)
}

func (r *Rucksack) Contains(i Item) bool {
	return util.Contains(i, r.Items())
}

func (r *Rucksack) FindDuplicateItems() []Item {
	duplicates := make([]Item, 0)

	for _, item := range r.Compartments[0] {
		if !util.Contains(item, duplicates) && util.Contains(item, r.Compartments[1]) {
			duplicates = append(duplicates, item)
		}
	}

	return duplicates
}

type Compartment []Item

type Item string

var itemPriorities = make([]string, 0)

func init() {
	for ch := 'a'; ch <= 'z'; ch++ {
		itemPriorities = append(itemPriorities, string(ch))
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		itemPriorities = append(itemPriorities, string(ch))
	}
}

func (i Item) Priority() int {
	for k, v := range itemPriorities {
		if string(i) == v {
			return k + 1
		}
	}
	return 0
}

const input = "day03/input.txt"

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rucksacks := make([]*Rucksack, 0)
	groups := make([]Group, 0)

	scanner := bufio.NewScanner(f)

	currentGroup := make(Group, 0)
	i := 0
	for scanner.Scan() {
		t := scanner.Text()
		rucksack := NewRucksack(t)
		rucksacks = append(rucksacks, rucksack)

		currentGroup = append(currentGroup, rucksack)

		i++
		if i == 3 {
			i = 0
			groups = append(groups, currentGroup)
			currentGroup = make(Group, 0)
		}
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(rucksacks), part2(groups))
}

func part1(rucksacks []*Rucksack) int {
	duplicateItems := make([]Item, 0)
	for _, r := range rucksacks {
		duplicateItems = append(duplicateItems, r.FindDuplicateItems()...)
	}

	return util.Sum(util.Map(duplicateItems, func(i Item) int {
		return i.Priority()
	}))
}

func part2(groups []Group) int {
	badges := util.Map(groups, func(g Group) Item {
		return g.FindBadge()
	})
	return util.Sum(util.Map(badges, func(i Item) int {
		return i.Priority()
	}))
}

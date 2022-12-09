package main

import (
	"bufio"
	"fmt"
	"github.com/TomWright/adventofcode/2022/util"
	"os"
	"strconv"
)

const input = "day08/input.txt"

type Map struct {
	Trees  []Tree
	Width  int
	Height int
}

func (m Map) String() string {
	res := ""
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			res += m.TreeAt(x, y).String()
		}
		res += "\n"
	}
	return res
}

func (m Map) IndexOf(x int, y int) int {
	return (m.Width * y) + x
}

func (m Map) TreeAt(x int, y int) Tree {
	return m.Trees[m.IndexOf(x, y)]
}

func (m Map) TreesLeft(x int, y int) []Tree {
	index := m.IndexOf(x, y)
	numTrees := x % m.Width
	if numTrees == 0 {
		return make([]Tree, 0)
	}
	return util.Reverse(m.Trees[index-numTrees : index])
}

func (m Map) TreesRight(x int, y int) []Tree {
	index := m.IndexOf(x, y)
	numTrees := m.Width - (x + 1)
	if numTrees == 0 {
		return make([]Tree, 0)
	}
	return m.Trees[index+1 : index+1+numTrees]
}

func (m Map) TreesAbove(x int, y int) []Tree {
	index := m.IndexOf(x, y) - m.Width
	res := make([]Tree, 0)
	for index >= 0 {
		res = append(res, m.Trees[index])
		index -= m.Width
	}
	return res
}

func (m Map) TreesBelow(x int, y int) []Tree {
	index := m.IndexOf(x, y) + m.Width
	res := make([]Tree, 0)
	for index < len(m.Trees) {
		res = append(res, m.Trees[index])
		index += m.Width
	}
	return res
}

func ViewDistance(tree Tree, trees []Tree) int {
	for k, t := range trees {
		if t.Height >= tree.Height {
			return k + 1
		}
	}
	return len(trees)
}

type Tree struct {
	Height int
	X      int
	Y      int
}

func (t Tree) String() string {
	return fmt.Sprint(t.Height)
}

func (t Tree) ScenicScore(m Map) int {
	return ViewDistance(t, m.TreesAbove(t.X, t.Y)) *
		ViewDistance(t, m.TreesLeft(t.X, t.Y)) *
		ViewDistance(t, m.TreesRight(t.X, t.Y)) *
		ViewDistance(t, m.TreesBelow(t.X, t.Y))
}

func (t Tree) Visible(m Map) bool {
	filterFn := func(other Tree) bool {
		return other.Height >= t.Height
	}
	return len(util.Filter(m.TreesLeft(t.X, t.Y), filterFn)) == 0 ||
		len(util.Filter(m.TreesRight(t.X, t.Y), filterFn)) == 0 ||
		len(util.Filter(m.TreesAbove(t.X, t.Y), filterFn)) == 0 ||
		len(util.Filter(m.TreesBelow(t.X, t.Y), filterFn)) == 0

}

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	m := Map{
		Trees:  make([]Tree, 0),
		Width:  0,
		Height: 0,
	}

	y := 0
	for scanner.Scan() {
		t := scanner.Text()

		for x, heightStr := range t {
			height, err := strconv.ParseInt(string(heightStr), 10, 64)
			if err != nil {
				panic(err)
			}

			m.Trees = append(m.Trees, Tree{
				Height: int(height),
				X:      x,
				Y:      y,
			})
		}

		if y == 0 {
			m.Width = len(m.Trees)
		}
		y++
	}
	m.Height = y

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(m), part2(m))
}

func part1(m Map) int {
	return len(util.Filter(m.Trees, func(t Tree) bool {
		return t.Visible(m)
	}))
}

func part2(m Map) int {
	max := 0
	for _, t := range m.Trees {
		cur := t.ScenicScore(m)
		if cur > max {
			max = cur
		}
	}
	return max
}

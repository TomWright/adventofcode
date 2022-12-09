package main

import (
	"bufio"
	"fmt"
	"github.com/TomWright/adventofcode/2022/util"
	"os"
	"strconv"
	"strings"
)

const input = "day09/input.txt"

type Direction string

var (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

type Move struct {
	Direction Direction
	Num       int
}

type XY struct {
	X int
	Y int
}

func (xy XY) DirectionTo(other XY) []Direction {
	res := make([]Direction, 0)

	if xy.IsNextTo(other) {
		return res
	}

	y := other.Y - xy.Y
	x := other.X - xy.X

	switch {
	case y > 0:
		res = append(res, Up)
	case y < 0:
		res = append(res, Down)
	}
	switch {
	case x > 0:
		res = append(res, Right)
	case x < 0:
		res = append(res, Left)
	}

	return res
}

func (xy XY) IsNextTo(other XY) bool {
	return other.X >= xy.X-1 && other.X <= xy.X+1 && other.Y >= xy.Y-1 && other.Y <= xy.Y+1
}

type Rope struct {
	Head          *Knot
	Tail          *Knot
	TailPositions []XY
}

func (g *Rope) MoveHead(m Move) {
	g.trackTail()
	for x := 0; x < m.Num; x++ {
		g.Head.Move([]Direction{m.Direction})
		g.trackTail()
	}
}

func (g *Rope) trackTail() {
	if !util.Contains(g.Tail.Position, g.TailPositions) {
		g.TailPositions = append(g.TailPositions, g.Tail.Position)
	}
}

type Knot struct {
	Name     string
	Position XY
	Next     *Knot
	Prev     *Knot
}

func (knot *Knot) Move(dirs []Direction) {
	for _, dir := range dirs {
		switch dir {
		case Up:
			knot.Position.Y++
		case Down:
			knot.Position.Y--
		case Left:
			knot.Position.X--
		case Right:
			knot.Position.X++
		}
	}

	if knot.Next != nil {
		knot.Next.Move(knot.Next.Position.DirectionTo(knot.Position))
	}
}

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	moves := make([]Move, 0)

	for scanner.Scan() {
		t := scanner.Text()

		parts := strings.Split(t, " ")
		num, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(err)
		}

		moves = append(moves, Move{
			Direction: Direction(parts[0]),
			Num:       int(num),
		})
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(moves), part2(moves))
}

// Create a linked list of knots up to numKnots deep.
func ropeWithKnots(numKnots int) *Rope {
	head := &Knot{
		Name:     "H",
		Position: XY{},
	}
	cur := head
	for x := 1; x < numKnots; x++ {
		next := &Knot{
			Name:     fmt.Sprint(x),
			Position: XY{},
			Prev:     cur,
		}
		cur.Next = next

		cur = next
	}
	r := &Rope{
		Head:          head,
		Tail:          cur,
		TailPositions: nil,
	}
	return r
}

func part1(moves []Move) int {
	r := ropeWithKnots(2)
	r.Tail.Name = "T"
	for _, m := range moves {
		r.MoveHead(m)
	}
	return len(r.TailPositions)
}

func part2(moves []Move) int {
	r := ropeWithKnots(10)
	for _, m := range moves {
		r.MoveHead(m)
	}
	return len(r.TailPositions)
}

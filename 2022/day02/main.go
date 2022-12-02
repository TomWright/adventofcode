package main

import (
	"bufio"
	"fmt"
	"github.com/TomWright/adventofcode/2022/util"
	"os"
	"strings"
)

type Move struct {
	Value  string
	Points int
}

func NewOpponentMove(move string) Move {
	switch move {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	default:
		panic("unknown opponent move: " + move)
	}
}

func NewResponseMove(move string) Move {
	switch move {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	default:
		panic("unknown response move: " + move)
	}
}

func NewResponseResult(expected string) Result {
	switch expected {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		panic("unknown response result: " + expected)
	}
}

type Result struct {
	Value  string
	Points int
}

var (
	Rock     = Move{Value: "Rock", Points: 1}
	Paper    = Move{Value: "Paper", Points: 2}
	Scissors = Move{Value: "Scissors", Points: 3}

	Win  = Result{Value: "Win", Points: 6}
	Lose = Result{Value: "Lose", Points: 0}
	Draw = Result{Value: "Draw", Points: 3}
)

func (m Move) Play(other Move) Result {
	if m == other {
		return Draw
	}
	if m == Rock && other == Scissors {
		return Win
	}
	if m == Scissors && other == Paper {
		return Win
	}
	if m == Paper && other == Rock {
		return Win
	}
	return Lose
}

const input = "day02/input.txt"

type Game struct {
	First  Move
	Second Move
}

func (g Game) Play() Result {
	// Second plays First because we are the responder.
	return g.Second.Play(g.First)
}

func (g Game) Points() int {
	return g.Play().Points + g.Second.Points
}

func NewGameFromString(moves string) Game {
	parts := strings.Split(moves, " ")
	return Game{
		First:  NewOpponentMove(parts[0]),
		Second: NewResponseMove(parts[1]),
	}
}

func NewExpectedResultGameFromString(moves string) Game {
	parts := strings.Split(moves, " ")
	firstMove := NewOpponentMove(parts[0])
	expectedResult := NewResponseResult(parts[1])

	for _, response := range []Move{
		Rock, Paper, Scissors,
	} {
		if response.Play(firstMove) == expectedResult {
			return Game{
				First:  firstMove,
				Second: response,
			}
		}
	}

	panic("could not find expected result")
}

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	part1Games := make([]Game, 0)
	part2Games := make([]Game, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		part1Games = append(part1Games, NewGameFromString(t))
		part2Games = append(part2Games, NewExpectedResultGameFromString(t))
	}

	part1 := util.Sum(util.Map(part1Games, func(g Game) int {
		return g.Points()
	}))

	part2 := util.Sum(util.Map(part2Games, func(g Game) int {
		return g.Points()
	}))

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1, part2)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const input = "day05/input.txt"

type Stacks []*Stack

func (s Stacks) Print() {
	for i, stack := range s {
		fmt.Printf("%d\t", i+1)
		for _, c := range stack.Crates {
			fmt.Printf("[%s] ", c)
		}
		fmt.Print("\n")
	}
	fmt.Print("----------\n")
}

func (s Stacks) Stack(num int) *Stack {
	return s[num-1]
}

func NewStack() *Stack {
	return &Stack{
		Crates: make([]Crate, 0),
	}
}

type Stack struct {
	Crates []Crate
}

func (s *Stack) AddCrateToBottom(c Crate) {
	s.Crates = append([]Crate{c}, s.Crates...)
}

func (s *Stack) AddCrate(c ...Crate) {
	s.Crates = append(s.Crates, c...)
}

func (s *Stack) GetCrate() Crate {
	return s.Crates[len(s.Crates)-1]
}

func (s *Stack) TakeCrate() Crate {
	return s.TakeCrates(1)[0]
}

func (s *Stack) TakeCrates(num int) []Crate {
	i := len(s.Crates) - num
	res := s.Crates[i:]
	s.Crates = s.Crates[:i]
	return res
}

type Crate string

var (
	moveRegex = regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
)

func NewMove(input string) Move {
	match := moveRegex.FindStringSubmatch(input)

	numCrates, _ := strconv.Atoi(match[1])
	from, _ := strconv.Atoi(match[2])
	to, _ := strconv.Atoi(match[3])

	return Move{
		NumCrates: numCrates,
		From:      from,
		To:        to,
	}
}

type Move struct {
	NumCrates int
	From      int
	To        int
}

func (m Move) String() string {
	return fmt.Sprintf("move %d from %d to %d", m.NumCrates, m.From, m.To)
}

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	part1Stacks := make(Stacks, 0)
	part2Stacks := make(Stacks, 0)
	moves := make([]Move, 0)

	processingMoves := false
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			processingMoves = true
			continue
		}

		if !processingMoves {

			reader := strings.NewReader(t)

			readNextPart := func() (string, bool) {
				part := ""
				for i := 0; i < 4; i++ {
					r, _, err := reader.ReadRune()
					if err == io.EOF {
						return part, part != ""
					}
					if err != nil {
						panic(err)
					}
					rStr := string(r)
					switch rStr {
					case " ":
						continue
					case "\n":
						continue
					default:
						part += rStr
					}

				}
				return part, true
			}

			stackNumber := 0
			for {
				stackNumber++

				part, found := readNextPart()
				if !found {
					break
				}

				if len(part1Stacks) < stackNumber {
					part1Stacks = append(part1Stacks, NewStack())
				}
				if len(part2Stacks) < stackNumber {
					part2Stacks = append(part2Stacks, NewStack())
				}

				if part == "" || part[0] != '[' {
					continue
				}
				part = strings.Trim(part, "[]")

				if part == "" {
					continue
				}
				c := Crate(part)
				part1Stacks.Stack(stackNumber).AddCrateToBottom(c)
				part2Stacks.Stack(stackNumber).AddCrateToBottom(c)
			}
		} else {
			moves = append(moves, NewMove(t))
		}
	}

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1(part1Stacks, moves), part2(part2Stacks, moves))
}

func part1(stacks Stacks, moves []Move) string {
	for _, m := range moves {
		for i := 0; i < m.NumCrates; i++ {
			c := stacks.Stack(m.From).TakeCrate()
			stacks.Stack(m.To).AddCrate(c)
		}
	}

	res := ""
	for _, s := range stacks {
		res += string(s.GetCrate())
	}
	return res
}

func part2(stacks Stacks, moves []Move) string {
	for _, m := range moves {
		c := stacks.Stack(m.From).TakeCrates(m.NumCrates)
		stacks.Stack(m.To).AddCrate(c...)
	}

	res := ""
	for _, s := range stacks {
		res += string(s.GetCrate())
	}
	return res
}

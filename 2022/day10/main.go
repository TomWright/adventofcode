package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const input = "day10/input.txt"

type CommandType string

var (
	NOOP CommandType = "noop"
	ADD  CommandType = "add"
)

func NewCommand(input string) *Command {
	p := strings.Split(input, " ")
	name := p[0]
	args := p[1:]
	if strings.HasPrefix(name, "add") {
		args = append([]string{strings.TrimPrefix(name, "add")}, args...)
		name = "add"
	}
	c := &Command{
		Name:   name,
		Args:   args,
		Status: Waiting,
	}
	return c
}

type Status string

const (
	Waiting    Status = "waiting"
	Processing Status = "processing"
	Done       Status = "done"
)

type Command struct {
	Name               string
	Args               []string
	remainingCycleTime int
	Status             Status
}

func (c *Command) String() string {
	switch c.Type() {
	case ADD:
		return fmt.Sprintf("add%s %s", c.Args[0], c.Args[1])
	default:
		return c.Name
	}
}

func (c *Command) Type() CommandType {
	return CommandType(c.Name)
}

func (c *Command) CycleTime() int {
	switch c.Type() {
	case NOOP:
		return 1
	case ADD:
		return 2
	default:
		return -1
	}
}

func (c *Command) Execute(registers *Registers) Status {
	if c.Status == Done {
		return c.Status
	}

	if c.Status == Waiting {
		c.Status = Processing
		c.remainingCycleTime = c.CycleTime()

		fmt.Println("Start execute", c)
	}

	c.remainingCycleTime--

	if c.remainingCycleTime > 0 {
		return c.Status
	}

	switch c.Type() {
	case NOOP:
	case ADD:
		reg := registers.Register(c.Args[0])
		num, err := strconv.Atoi(c.Args[1])
		if err != nil {
			panic(err)
		}
		reg.value += num
	}

	c.Status = Done
	return c.Status
}

func NewRegisters() *Registers {
	return &Registers{registers: map[string]*Register{
		"x": NewRegister(),
	}}
}

type Registers struct {
	registers map[string]*Register
}

func (r Registers) String() string {
	res := ""
	for k, re := range r.registers {
		res += fmt.Sprintf("register %s: %d\n", k, re.value)
	}
	return strings.TrimSpace(res)
}

func (r Registers) Register(name string) *Register {
	if _, ok := r.registers[name]; !ok {
		r.registers[name] = NewRegister()
	}
	return r.registers[name]
}

func NewRegister() *Register {
	return &Register{value: 1}
}

type Register struct {
	value int
}

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	commands := make([]*Command, 0)

	for scanner.Scan() {
		t := scanner.Text()

		commands = append(commands, NewCommand(t))
	}

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(commands), part2(commands))
}

func runAllCommands(commands []*Command, registers *Registers, preExecFn func(cycle int, registers *Registers)) {
	commandsCh := make(chan *Command, len(commands))
	for _, c := range commands {
		commandsCh <- c
	}
	close(commandsCh)

	var command *Command = nil
	cycle := 0
	for {
		if command == nil || command.Status == Done {
			newCommand, ok := <-commandsCh
			if !ok {
				return
			}
			command = newCommand
		}

		fmt.Printf("----- Start cycle %d -----\n", cycle+1)
		fmt.Println("Start", *registers)

		if preExecFn != nil {
			preExecFn(cycle+1, registers)
		}

		command.Execute(registers)

		fmt.Println("End", *registers)
		fmt.Printf("----- End cycle %d -----\n", cycle+1)
		cycle++
	}
}

func part1(commands []*Command) int {
	registers := NewRegisters()

	res := 0

	preExecFn := func(cycle int, registers *Registers) {
		if (cycle-20)%40 == 0 {
			signalStrength := cycle * registers.Register("x").value
			res += signalStrength
		}
	}

	runAllCommands(commands, registers, preExecFn)

	return res
}

func part2(commands []*Command) int {
	return 0
}

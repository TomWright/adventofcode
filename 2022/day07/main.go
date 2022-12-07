package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const input = "day07/input.txt"

func NewCommand(input string) *Command {
	parts := strings.Split(input, " ")
	return &Command{
		Command: parts[1],
		Args:    parts[2:],
		Output:  nil,
	}
}

type Command struct {
	Command string
	Args    []string
	Output  []string
}

type Commands []Command

func (commands Commands) Filesystem() *Node {
	root := &Node{
		Name:     "/",
		Type:     Directory,
		Children: make([]*Node, 0),
		Size:     0,
		Parent:   nil,
	}
	var cwd *Node

	for _, c := range commands {
		switch c.Command {
		case "cd":
			switch c.Args[0] {
			case "/":
				cwd = root
			case "..":
				cwd = cwd.Parent
			default:
				set := false
				for _, ch := range cwd.Children {
					if ch.Type != Directory {
						continue
					}
					if ch.Name == c.Args[0] {
						set = true
						cwd = ch
						break
					}
				}
				if !set {
					panic("could not cd")
				}
			}
		case "ls":
			cwd.Children = make([]*Node, 0)
			for _, o := range c.Output {
				if strings.HasPrefix(o, "dir ") {
					cwd.Children = append(cwd.Children, &Node{
						Name:     strings.TrimPrefix(o, "dir "),
						Type:     Directory,
						Children: make([]*Node, 0),
						Size:     0,
						Parent:   cwd,
					})
					continue
				}
				parts := strings.Split(o, " ")
				size, err := strconv.ParseInt(parts[0], 10, 64)
				if err != nil {
					panic(err)
				}
				cwd.Children = append(cwd.Children, &Node{
					Name:     parts[1],
					Type:     File,
					Children: nil,
					Size:     size,
					Parent:   cwd,
				})
			}
			cwd.Size = int64(len(cwd.Children))
		default:
			panic("unhandled command " + c.Command)
		}
	}

	return root
}

type NodeType string

const (
	File      NodeType = "file"
	Directory NodeType = "dir"
)

type Node struct {
	Name     string
	Type     NodeType
	Children []*Node
	Size     int64
	Parent   *Node
}

func (n *Node) DirSize() int64 {
	size := int64(0)
	for _, c := range n.Children {
		switch c.Type {
		case File:
			size += c.Size
		case Directory:
			size += c.DirSize()
		}
	}
	return size
}

func (n *Node) Filter(fn func(*Node) bool) []*Node {
	res := make([]*Node, 0)
	if fn(n) {
		res = append(res, n)
	}
	if n.Type == Directory {
		for _, c := range n.Children {
			res = append(res, c.Filter(fn)...)
		}
	}
	return res
}

func (n *Node) Print() {
	n.print(0)
}

func (n *Node) print(level int) {
	indent := strings.Repeat("  ", level)

	switch n.Type {
	case Directory:
		fmt.Printf("%s- %s (%s)\n", indent, n.Name, n.Type)
		for _, c := range n.Children {
			c.print(level + 1)
		}
	case File:
		fmt.Printf("%s- %s (%s, size=%d)\n", indent, n.Name, n.Type, n.Size)
		for _, c := range n.Children {
			c.print(level + 1)
		}
	}
}

func main() {
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	commands := make(Commands, 0)

	var command *Command

	for scanner.Scan() {
		t := scanner.Text()
		if strings.HasPrefix(t, "$ ") {
			if command != nil {
				commands = append(commands, *command)
			}
			command = NewCommand(t)
		} else {
			command.Output = append(command.Output, t)
		}
	}
	commands = append(commands, *command)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(commands), part2(commands))
}

func part1(commands Commands) int64 {
	fs := commands.Filesystem()

	const limit = int64(100000)
	found := fs.Filter(func(node *Node) bool {
		return node.Name != "/" && node.Type == Directory && node.DirSize() <= limit
	})

	sum := int64(0)
	for _, f := range found {
		sum += f.DirSize()
	}
	return sum
}

func part2(commands Commands) int64 {
	fs := commands.Filesystem()

	const totalSpace = int64(70000000)
	const requiredSpace = int64(30000000)
	usedSpace := fs.DirSize()
	unusedSpace := totalSpace - usedSpace
	missingSpace := requiredSpace - unusedSpace

	found := fs.Filter(func(node *Node) bool {
		return node.Name != "/" && node.Type == Directory && node.DirSize() >= missingSpace
	})

	sort.Slice(found, func(i, j int) bool {
		return found[i].DirSize() < found[j].DirSize()
	})

	return found[0].DirSize()
}

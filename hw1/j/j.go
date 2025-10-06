package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Namespace struct {
	storage map[string]GenericList
}

func NewNamespace() *Namespace {
	return &Namespace{
		storage: make(map[string]GenericList),
	}
}

func (n *Namespace) Get(name string) GenericList {
	return n.storage[name]
}

func (n *Namespace) Set(name string, value GenericList) {
	n.storage[name] = value
}

type GenericList interface {
	Get(i int) int
	Set(i int, value int)
	Sublist(from, to int) GenericList
	Length() int
	Add(value int)
}

type List struct {
	data       []int
	isRootList bool
}

func NewList(data []int) *List {
	return &List{
		data:       data,
		isRootList: true,
	}
}

func (l *List) Get(i int) int {
	return l.data[i]
}

func (l *List) Set(i int, value int) {
	l.data[i] = value
}

func (l *List) Add(value int) {
	l.data = append(l.data, value)
}

func (l *List) Sublist(from, to int) GenericList {
	return NewSublist(l, from, to)
}

func (l *List) Length() int {
	return len(l.data)
}

type Sublist struct {
	parent     GenericList
	start      int
	end        int
	isRootList bool
}

func NewSublist(parent GenericList, start, end int) *Sublist {
	return &Sublist{
		parent:     parent,
		start:      start,
		end:        end,
		isRootList: false,
	}
}

func (s *Sublist) Get(i int) int {
	return s.parent.Get(s.start + i)
}

func (s *Sublist) Set(i int, value int) {
	s.parent.Set(s.start+i, value)
}

func (s *Sublist) Sublist(from, to int) GenericList {
	return NewSublist(s.parent, s.start+from, s.start+to)
}

func (s *Sublist) Length() int {
	return s.end - s.start + 1
}

func (s *Sublist) Add(value int) {
	panic("Add is only available for root lists")
}

func executeCommand(line string, memory *Namespace, writer *bufio.Writer) {
	line = strings.TrimSpace(line)

	if strings.HasPrefix(line, "List ") && strings.Contains(line, "= new List(") {
		parts := strings.Split(line, "=")
		name := strings.TrimSpace(strings.Fields(parts[0])[1])
		right := strings.TrimSpace(parts[1])

		valuesStr := right[len("new List(") : len(right)-1]
		var values []int

		if valuesStr != "" {
			valueStrs := strings.Split(valuesStr, ",")
			values = make([]int, len(valueStrs))
			for i, str := range valueStrs {
				val, _ := strconv.Atoi(strings.TrimSpace(str))
				values[i] = val
			}
		}

		memory.Set(name, NewList(values))

	} else if strings.HasPrefix(line, "List ") && strings.Contains(line, ".subList(") {
		parts := strings.Split(line, "=")
		newName := strings.TrimSpace(strings.Fields(parts[0])[1])
		right := strings.TrimSpace(parts[1])

		sourceName := strings.Split(right, ".subList(")[0]
		source := memory.Get(sourceName)

		argsStr := right[strings.Index(right, "(")+1 : len(right)-1]
		args := strings.Split(argsStr, ",")
		start, _ := strconv.Atoi(strings.TrimSpace(args[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(args[1]))

		start0based := start - 1
		end0based := end - 1

		memory.Set(newName, source.Sublist(start0based, end0based))

	} else if strings.Contains(line, ".get(") {
		dotIndex := strings.Index(line, ".get(")
		name := line[:dotIndex]
		argsStr := line[dotIndex+len(".get(") : len(line)-1]

		lst := memory.Get(name)
		index, _ := strconv.Atoi(strings.TrimSpace(argsStr))
		index0based := index - 1

		writer.WriteString(strconv.Itoa(lst.Get(index0based)))
		writer.WriteByte('\n')

	} else if strings.Contains(line, ".set(") {
		dotIndex := strings.Index(line, ".set(")
		name := line[:dotIndex]
		argsStr := line[dotIndex+len(".set(") : len(line)-1]

		lst := memory.Get(name)
		args := strings.Split(argsStr, ",")
		index, _ := strconv.Atoi(strings.TrimSpace(args[0]))
		value, _ := strconv.Atoi(strings.TrimSpace(args[1]))

		index0based := index - 1
		lst.Set(index0based, value)

	} else if strings.Contains(line, ".add(") {
		dotIndex := strings.Index(line, ".add(")
		name := line[:dotIndex]
		argsStr := line[dotIndex+len(".add(") : len(line)-1]

		lst := memory.Get(name)
		value, _ := strconv.Atoi(strings.TrimSpace(argsStr))
		lst.Add(value)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<23)
	writer := bufio.NewWriterSize(os.Stdout, 1<<23)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	memory := NewNamespace()

	for range n {
		line, _ := reader.ReadString('\n')
		executeCommand(line, memory, writer)
	}
}

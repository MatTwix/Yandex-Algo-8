package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Disk struct {
	memory        []int
	firstFreeCell int
}

func NewDisk(size int) *Disk {
	return &Disk{memory: make([]int, size)}
}

func (d *Disk) Allocate(size int) int {
	addr := d.firstFreeCell
	d.firstFreeCell += size
	return addr
}

type Sequence interface {
	Get(idx int) int
	Set(idx, x int)
	StartPtr() int
	Disk() *Disk
}

type List struct {
	name     string
	startPtr int
	len      int
	capacity int
	disk     *Disk
}

func nextTwoPower(n int) int {
	if n <= 1 {
		return 1
	}
	p := 1
	for p < n {
		p <<= 1
	}
	return p
}

func NewList(disk *Disk, name string, elems ...int) *List {
	length := len(elems)
	cap := nextTwoPower(length)

	start := disk.Allocate(cap)
	for idx, el := range elems {
		disk.memory[start+idx] = el
	}

	return &List{
		name:     name,
		startPtr: start,
		len:      length,
		capacity: cap,
		disk:     disk,
	}
}

func (l *List) Add(x int) {
	if l.len == l.capacity {
		newCapacity := l.capacity * 2
		newStart := l.disk.Allocate(newCapacity)
		copy(l.disk.memory[newStart:], l.disk.memory[l.startPtr:l.startPtr+l.len])
		l.startPtr = newStart
		l.capacity = newCapacity
	}

	l.disk.memory[l.startPtr+l.len] = x
	l.len++
}

func (l *List) Set(idx, x int) {
	l.disk.memory[l.startPtr+idx] = x
}

func (l *List) Get(idx int) int {
	return l.disk.memory[l.startPtr+idx]
}

func (l *List) StartPtr() int {
	return l.startPtr
}

func (l *List) Disk() *Disk {
	return l.disk
}

type Sublist struct {
	name     string
	parent   Sequence
	startIdx int
	endIdx   int
	len      int
}

func NewSublist(name string, parent Sequence, startIdx, endIdx int) *Sublist {
	startIdx--
	endIdx--

	return &Sublist{
		name:     name,
		parent:   parent,
		startIdx: startIdx,
		endIdx:   endIdx,
		len:      endIdx - startIdx + 1,
	}
}

func (s *Sublist) Set(idx, x int) {
	basePrt := s.parent.StartPtr()
	disk := s.parent.Disk()
	disk.memory[basePrt+s.startIdx+idx] = x
}

func (s *Sublist) Get(idx int) int {
	basePrt := s.parent.StartPtr()
	disk := s.parent.Disk()
	return disk.memory[basePrt+s.startIdx+idx]
}

func (s *Sublist) StartPtr() int {
	return s.parent.StartPtr() + s.startIdx
}

func (s *Sublist) Disk() *Disk {
	return s.parent.Disk()
}

type Interpreter struct {
	disk      *Disk
	sequences map[string]Sequence
	writer    *bufio.Writer
}

func NewInterpreter(disk *Disk, writer *bufio.Writer) *Interpreter {
	return &Interpreter{
		disk:      disk,
		sequences: make(map[string]Sequence),
		writer:    writer,
	}
}

func (i *Interpreter) AddSequence(name string, seq Sequence) {
	i.sequences[name] = seq
}

func (i *Interpreter) GetSequence(name string) Sequence {
	return i.sequences[name]
}

func (i *Interpreter) ExecuteCommand(line string) {
	if strings.HasPrefix(line, "List") {
		parts := strings.Split(line, "=")
		name := strings.TrimSpace(strings.Fields(parts[0])[1])
		right := strings.TrimSpace(parts[1])

		if strings.HasPrefix(right, "new List") {
			content := right[len("new List(") : len(right)-1]
			nums := strings.Split(content, ",")

			elems := make([]int, len(nums))

			for idx, s := range nums {
				elems[idx], _ = strconv.Atoi(s)
			}

			lst := NewList(i.disk, name, elems...)
			i.AddSequence(name, lst)
		} else if strings.Contains(right, ".subList") {
			left := strings.Split(right, ".subList")[0]

			seq := i.GetSequence(left)

			content := right[strings.Index(right, "(")+1 : len(right)-1]
			bounds := strings.Split(content, ",")

			from, _ := strconv.Atoi(bounds[0])
			to, _ := strconv.Atoi(bounds[1])

			sub := NewSublist(name, seq, from, to)
			i.AddSequence(name, sub)
		}
	} else if strings.Contains(line, ".set") {
		dot := strings.Index(line, ".")
		name := line[:dot]
		seq := i.GetSequence(name)
		args := line[dot+len(".set(") : len(line)-1]
		idxVal := strings.Split(args, ",")
		idx, _ := strconv.Atoi(idxVal[0])
		val, _ := strconv.Atoi(idxVal[1])
		seq.Set(idx-1, val)
	} else if strings.Contains(line, ".add") {
		parts := strings.Split(line, ".")
		name := parts[0]

		seq := i.GetSequence(name)

		val, _ := strconv.Atoi(parts[1][len("add(") : len(parts[1])-1])
		if l, ok := seq.(*List); ok {
			l.Add(val)
		}
	} else if strings.Contains(line, ".get") {
		parts := strings.Split(line, ".")
		name := parts[0]

		seq := i.GetSequence(name)

		idx, _ := strconv.Atoi(parts[1][len("get(") : len(parts[1])-1])
		i.writer.WriteString(strconv.Itoa(seq.Get(idx - 1)))
		i.writer.WriteByte('\n')
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	commands := []string{}

	for range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		commands = append(commands, line)
	}

	disk := NewDisk(2_000_000)

	interp := NewInterpreter(disk, writer)

	for _, cmd := range commands {
		interp.ExecuteCommand(cmd)
	}
}

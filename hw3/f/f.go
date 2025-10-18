package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	parentLine, _ := reader.ReadString('\n')
	parentParts := strings.Fields(parentLine)
	parentNodes := make([]int, n)
	for i := range n {
		parentNodes[i], _ = strconv.Atoi(parentParts[i])
	}

	tree := make([][]int, n)
	root := -1
	for i, p := range parentNodes {
		if p == 0 {
			root = i
		} else {
			tree[p-1] = append(tree[p-1], i)
		}
	}

	timecount := make([][2]int, n)
	for i := range n {
		timecount[i][0], timecount[i][1] = -1, -1
	}
	enter := make([]bool, n)
	counter := 0

	type pair struct {
		node, state int
	}

	stack := []pair{{root, 0}}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := top.node
		state := top.state

		if !enter[node] {
			timecount[node][0] = counter
			counter++
			enter[node] = true
			stack = append(stack, pair{node, 1})
			for i := len(tree[node]) - 1; i >= 0; i-- {
				child := tree[node][i]
				stack = append(stack, pair{child, 0})
			}
		} else if state == 1 {
			timecount[node][1] = counter
			counter++
		}
	}

	mLine, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mLine))

	for range m {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		a--
		b--

		if timecount[a][0] <= timecount[b][0] &&
			timecount[b][0] <= timecount[b][1] &&
			timecount[b][1] <= timecount[a][1] {
			writer.WriteString("1\n")
		} else {
			writer.WriteString("0\n")
		}
	}
}

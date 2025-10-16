package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type node struct {
	nearestLeaf int
	neigh       []int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	graph := make([]node, n)

	for range n - 1 {
		argLine, _ := reader.ReadString('\n')
		argParts := strings.Fields(argLine)
		a, _ := strconv.Atoi(argParts[0])
		b, _ := strconv.Atoi(argParts[1])

		a--
		b--

		graph[a].neigh = append(graph[a].neigh, b)
		graph[b].neigh = append(graph[b].neigh, a)
	}

	mn := 1_000_000_000

	dfs(0, -1, &mn, graph)

	if mn == 1_000_000_000 {
		mn = n - 1
	}

	writer.WriteString(strconv.Itoa(mn))
	writer.WriteByte('\n')
}

func dfs(idx, from int, mn *int, graph []node) {
	if len(graph[idx].neigh) > 2 || from == -1 && len(graph[idx].neigh) >= 2 {
		fNLD, sNLD := 1_000_000_000, 1_000_000_000

		for _, i := range graph[idx].neigh {
			if i == from {
				continue
			}
			dfs(i, idx, mn, graph)

			dist := graph[i].nearestLeaf + 1

			switch {
			case dist <= fNLD:
				sNLD = fNLD
				fNLD = dist
			case dist < sNLD:
				sNLD = dist
			}
		}

		*mn = min(*mn, fNLD+sNLD)

		graph[idx].nearestLeaf = fNLD
	} else if len(graph[idx].neigh) == 2 || from == -1 && len(graph[idx].neigh) == 1 {
		child := graph[idx].neigh[0]
		if from == child {
			child = graph[idx].neigh[1]
		}

		dfs(child, idx, mn, graph)

		graph[idx].nearestLeaf = graph[child].nearestLeaf + 1
	} else {
		graph[idx].nearestLeaf = 0
	}
}

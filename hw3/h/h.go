package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type node struct {
	people   int64
	neigh    []int
	sumNeigh int64
	parent   int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	city := make([]node, n)

	argLine, _ := reader.ReadString('\n')
	parts := strings.Fields(argLine)

	for i := range parts {
		city[i].people, _ = strconv.ParseInt(parts[i], 10, 64)
	}

	for range n - 1 {
		connLine, _ := reader.ReadString('\n')
		connParts := strings.Fields(connLine)
		a, _ := strconv.Atoi(connParts[0])
		b, _ := strconv.Atoi(connParts[1])

		a--
		b--

		city[a].neigh = append(city[a].neigh, b)
		city[b].neigh = append(city[b].neigh, a)
	}

	totalSum := dfs(city, 0, -1)

	ansVal := int64(math.MaxInt64)
	ans := -1

	for i, v := range city {
		mxChild := int64(-1)
		// println("children:", i)
		for _, c := range v.neigh {
			if c == city[i].parent {
				continue
			}
			mxChild = max(mxChild, city[c].sumNeigh)
			// println(c, city[c].sumNeigh)
		}

		candidate := max(v.people, totalSum-v.sumNeigh, mxChild)
		if candidate < ansVal {
			ansVal = candidate
			ans = i
		}

		// println(i, v.people, totalSum-v.sumNeigh, mxChild)
	}

	// fmt.Printf("city: %v\nans: %d, ansVal: %d\n", city, ans, ansVal)

	writer.WriteString(strconv.Itoa(ans + 1))
	writer.WriteByte('\n')
}

func dfs(city []node, idx int, parent int) (sm int64) {
	sm = city[idx].people

	for _, i := range city[idx].neigh {
		if i == parent {
			continue
		}

		city[i].parent = idx
		sm += dfs(city, i, idx)
	}

	city[idx].sumNeigh = sm

	return sm
}

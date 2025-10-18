package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(u int, totalEdicts *int64, adj [][]int, balances []int64) int64 {
	var smChildren int64 = 0

	for _, v := range adj[u] {
		smChildren += dfs(v, totalEdicts, adj, balances)
	}

	sU := -balances[u]

	eU := sU - smChildren

	if eU > 0 {
		*totalEdicts += eU
	} else {
		*totalEdicts -= eU
	}

	return sU
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	adj := make([][]int, n)
	balances := make([]int64, n)

	for i := 1; i < n; i++ {
		argLine, _ := reader.ReadString('\n')
		p, _ := strconv.Atoi(strings.TrimSpace(argLine))
		adj[p] = append(adj[p], i)
	}

	partLine, _ := reader.ReadString('\n')
	parts := strings.Fields(partLine)

	for i, v := range parts {
		vI, _ := strconv.Atoi(v)
		balances[i] = int64(vI)
	}

	var totalEdicts int64

	dfs(0, &totalEdicts, adj, balances)

	fmt.Println(totalEdicts)
}

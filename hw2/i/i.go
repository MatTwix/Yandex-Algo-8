package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	n, m      int
	field, dp [][]int
	dx        = []int{1, -1, 0, 0}
	dy        = []int{0, 0, 1, -1}
)

func dfs(i, j int) int {
	if dp[i][j] != -1 {
		return dp[i][j]
	}

	res := 1
	for dir := range 4 {
		ni := i + dx[dir]
		nj := j + dy[dir]
		if ni >= 0 && ni < n && nj >= 0 && nj < m && field[ni][nj] == field[i][j]+1 {
			res = max(res, 1+dfs(ni, nj))
		}
	}
	dp[i][j] = res

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ = strconv.Atoi(parts[0])
	m, _ = strconv.Atoi(parts[1])

	field = make([][]int, n)
	dp = make([][]int, n)

	for i := range n {
		field[i] = make([]int, m)
		dp[i] = make([]int, m)
		for j := range m {
			dp[i][j] = -1
		}

		aLine, _ := reader.ReadString('\n')
		aParts := strings.Fields(aLine)
		for j := range m {
			field[i][j], _ = strconv.Atoi(aParts[j])
		}
	}

	ans := 0
	for i := range n {
		for j := range m {
			ans = max(ans, dfs(i, j))
		}
	}

	// for _, row := range field {
	// 	for _, val := range row {
	// 		print(val, " ")
	// 	}
	// 	println()
	// }

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

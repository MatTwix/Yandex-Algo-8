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

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for k := 0; k <= n; k++ {
		dp[0][k] = 1
	}

	for i := 1; i <= n; i++ {
		for k := 1; k <= n; k++ {
			dp[i][k] = dp[i][k-1]
			if i >= k {
				dp[i][k] += dp[i-k][k-1]
			}
		}
	}

	writer.WriteString(strconv.Itoa(dp[n][n]))
	writer.WriteByte('\n')
}

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

	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	var ans int

	switch n {
	case 0:
		ans = 1
	case 1:
		ans = 1
	case 2:
		ans = 2
	default:
		var dp []int
		dp = make([]int, n+1)
		dp[0], dp[1], dp[2] = 1, 1, 2
		for i := 3; i < n+1; i++ {
			dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
		}

		ans = dp[n]
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

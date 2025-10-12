package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	start, end, weight float64
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	if n == 0 {
		writer.WriteString("0")
		writer.WriteByte('\n')
		return
	}

	intervals := make([]interval, n)
	for i := range n {
		iLine, _ := reader.ReadString('\n')
		iParts := strings.Fields(iLine)
		var itr interval
		itr.start, _ = strconv.ParseFloat(iParts[0], 32)
		itr.end, _ = strconv.ParseFloat(iParts[1], 32)
		itr.weight, _ = strconv.ParseFloat(iParts[2], 32)

		intervals[i] = itr
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})

	p := make([]int, n)
	for i := range n {
		l, r := 0, i-1
		p[i] = -1
		for l <= r {
			m := (l + r) / 2
			if intervals[m].end <= intervals[i].start {
				p[i] = m
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	dp := make([]float64, n)
	for i := range n {
		mx := intervals[i].weight
		if p[i] != -1 {
			mx += dp[p[i]]
		}
		if i > 0 {
			dp[i] = max(dp[i-1], mx)
		} else {
			dp[i] = mx
		}
	}

	fmt.Fprintf(writer, "%.4f\n", dp[n-1])
}

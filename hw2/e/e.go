package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type pair struct {
	val  int
	prev int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	aLine, _ := reader.ReadString('\n')
	aParts := strings.Split(strings.TrimSpace(aLine), " ")

	a := make([]int, len(aParts))
	for i, s := range aParts {
		a[i], _ = strconv.Atoi(s)
	}

	protections := make([]int, n)
	for i := range n - k + 1 {
		sm := 0
		mn := a[i]
		for j := range k {
			sm += a[i+j]
			mn = min(mn, a[i+j])
		}
		protections[i] = sm * mn
	}

	dp := make([]pair, n+1)
	for i := range dp {
		dp[i].prev = -1
	}

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]

		if i >= k {
			candidate := dp[i-k].val + protections[i-k]
			if candidate > dp[i].val {
				dp[i].val = candidate
				dp[i].prev = i - k
			}
		}
	}

	ansPos := []string{}
	for i := n; i > 0; {
		if dp[i].prev != -1 && dp[i].val == dp[dp[i].prev].val+protections[dp[i].prev] {
			ansPos = append(ansPos, strconv.Itoa((dp[i].prev + 1)))
			i = dp[i].prev
		} else {
			i--
		}
	}

	slices.Reverse(ansPos)

	writer.WriteString(strconv.Itoa(len(ansPos)))
	writer.WriteByte('\n')
	writer.WriteString(strings.Join(ansPos, " "))
	writer.WriteByte('\n')
}

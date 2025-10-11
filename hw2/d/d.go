package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type pair struct {
	exists bool
	prev   int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	dict := make(map[string]bool, n)
	words := make([]string, n)

	for i := range n {
		wLine, _ := reader.ReadString('\n')
		w := strings.TrimSpace(wLine)
		words[i] = w
		dict[w] = true
	}

	dp := make([]pair, len(s)+1)
	for i := range dp {
		dp[i].prev = -1
	}

	dp[0].exists = true

	for i := 1; i <= len(s); i++ {
		for l := 1; l <= 20 && l <= i; l++ {
			if !dp[i-l].exists {
				continue
			}

			substr := s[i-l : i]
			if dict[substr] {
				dp[i].exists = true
				dp[i].prev = i - l
				break
			}
		}
	}

	parts := []string{}
	for i := len(s); i > 0; i = dp[i].prev {
		parts = append(parts, s[dp[i].prev:i])
	}

	slices.Reverse(parts)

	writer.WriteString(strings.Join(parts, " "))
	writer.WriteByte('\n')
}

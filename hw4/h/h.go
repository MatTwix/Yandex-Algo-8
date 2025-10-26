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

	a := make([]int, n)
	aLine, _ := reader.ReadString('\n')
	aParts := strings.Fields(aLine)

	for i, v := range aParts {
		a[i], _ = strconv.Atoi(v)
	}

	prefixSums := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		prefixSums[i] = prefixSums[i-1] + int64(a[i-1])
	}

	total := int64(0)
	for j := range n {
		l := j + 1
		r := min(j+a[j]-1, n-1)

		if l <= r {
			total += prefixSums[r+1] - prefixSums[l]
		}
	}

	writer.WriteString(strconv.FormatInt(total, 10))
	writer.WriteByte('\n')
}

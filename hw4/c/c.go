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

	nxLine, _ := reader.ReadString('\n')
	nxParts := strings.Fields(nxLine)
	n, _ := strconv.Atoi(nxParts[0])
	x, _ := strconv.Atoi(nxParts[1])

	cLine, _ := reader.ReadString('\n')
	cParts := strings.Fields(cLine)

	mLine, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mLine))

	workingArr := make([]int, n+m)

	for i := range n {
		c, _ := strconv.Atoi(cParts[i])
		if c >= x {
			workingArr[i] = 1
		}
	}

	prefixSums := make([]int, n+m+1)

	for i := 0; i < n+m; i++ {
		prefixSums[i+1] = prefixSums[i] + workingArr[i]
	}

	head := 0
	tail := n

	for range m {
		opLine, _ := reader.ReadString('\n')
		opParts := strings.Fields(opLine)
		op, _ := strconv.Atoi(opParts[0])

		switch op {
		case 1:
			a, _ := strconv.Atoi(opParts[1])

			if a >= x {
				workingArr[tail] = 1
			}
			prefixSums[tail+1] = prefixSums[tail] + workingArr[tail]
			tail++
		case 2:
			head++
		case 3:
			k, _ := strconv.Atoi(opParts[1])

			writer.WriteString(strconv.Itoa((prefixSums[head+k] - prefixSums[head])))
			writer.WriteByte('\n')
		}
	}
}

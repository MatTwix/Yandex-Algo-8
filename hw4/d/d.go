package main

import (
	"bufio"
	"fmt"
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

	agrsLine, _ := reader.ReadString('\n')
	argsField := strings.Fields(agrsLine)

	tables := make([]int, n)

	for i, v := range argsField {
		tables[i], _ = strconv.Atoi(v)
	}

	lSum := tables[0]
	rSum := tables[n-1]
	l := 0
	r := n - 1
	ans := module(lSum - rSum)
	ansL, ansR := 1, n

	for l+1 < r {
		if lSum <= rSum {
			l++
			lSum += tables[l]
		} else {
			r--
			rSum += tables[r]
		}

		if ans > module(lSum-rSum) {
			ansL = l + 1
			ansR = r + 1
			ans = module(lSum - rSum)
		}
	}

	fmt.Fprintf(writer, "%d %d %d\n", ans, ansL, ansR)
}

func module(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

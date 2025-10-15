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

	sLine, _ := reader.ReadString('\n')
	sParts := strings.Fields(sLine)

	a, _ := strconv.Atoi(sParts[0])
	b, _ := strconv.Atoi(sParts[1])
	s, _ := strconv.Atoi(sParts[2])

	ans := binSearch(max(a, b), 10_000_100, []int{a, b, s})
	if s != (ans-a)*(ans-b) {
		ans = -1
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func binSearch(l, r int, params []int) int {
	for l < r {
		m := (l + r) / 2
		if params[2] <= (m-params[0])*(m-params[1]) {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

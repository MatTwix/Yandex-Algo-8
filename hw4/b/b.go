package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type tax struct {
	power   int
	payment int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	taxes := make([]tax, n)

	for i := range n {
		argLine, _ := reader.ReadString('\n')
		parts := strings.Fields(argLine)
		taxes[i].power, _ = strconv.Atoi(parts[0])
		taxes[i].payment, _ = strconv.Atoi(parts[1])
	}

	mLine, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mLine))

	for range m {
		argLine, _ := reader.ReadString('\n')
		arg, _ := strconv.Atoi(strings.TrimSpace(argLine))
		ans := arg * findTax(arg, taxes)

		writer.WriteString(strconv.Itoa(ans))
		writer.WriteByte('\n')
	}
}

func findTax(power int, taxes []tax) int {
	l := 0
	r := len(taxes) - 1
	ans := -1

	for l <= r {
		m := l + (r-l)/2
		if taxes[m].power < power {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return taxes[ans].payment
}

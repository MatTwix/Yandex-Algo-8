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
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	table := make([][]rune, n)
	for i := range n {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		table[i] = []rune(line)
	}

	rules := map[rune]int{
		'+': 1,
		'-': -1,
	}

	rowsSum := make([]int, n)
	rowsQs := make([]int, n)

	colsSum := make([]int, m)
	colsQs := make([]int, m)

	for i := range n {
		for j := range m {
			if table[i][j] == '?' {
				rowsQs[i]++
			} else {
				rowsSum[i] += rules[table[i][j]]
			}
		}
	}

	for j := range m {
		for i := range n {
			if table[i][j] == '?' {
				colsQs[j]++
			} else {
				colsSum[j] += rules[table[i][j]]
			}
		}
	}

	res := -10000000
	for i := range n {
		for j := range m {
			mxRow, mnCol := rowsSum[i]+rowsQs[i], colsSum[j]-colsQs[j]
			if table[i][j] == '?' {
				mxRow--
				mnCol++
			}

			res = max(res, mxRow-mnCol)
		}
	}

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}

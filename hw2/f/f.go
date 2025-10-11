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

	field := make([][]rune, n)

	table := make([][]int, n+1)
	for i := range n + 1 {
		table[i] = []int{-1, 0, 0, 0, -1}
	}

	for i := range n {
		fLine, _ := reader.ReadString('\n')
		field[i] = []rune(strings.TrimSpace(fLine))
	}

	ans := 0

	for i := 1; i < n+1; i++ {
		canGo := false
		for j := 1; j <= 3; j++ {
			v := max(table[i-1][j-1], table[i-1][j], table[i-1][j+1])
			localCanGo := v > -1 && field[i-1][j-1] != 'W'

			table[i][j] = v

			if !localCanGo {
				table[i][j] = -1
			} else if field[i-1][j-1] == 'C' {
				table[i][j]++
			}

			ans = max(ans, table[i][j])

			canGo = canGo || localCanGo
		}
		if !canGo {
			break
		}
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

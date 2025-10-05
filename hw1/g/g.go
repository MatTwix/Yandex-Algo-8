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

	field := make([][]rune, n)
	for i := range field {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		field[i] = []rune(line)
	}

	dirs := [4][2]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
	}

	for i := range n {
		for j := range m {
			if field[i][j] == '.' {
				continue
			}
			for _, d := range dirs {
				cnt := 1
				for step := 1; step < 5; step++ {
					ni := i + d[0]*step
					nj := j + d[1]*step
					if ni < 0 || ni >= n || nj < 0 || nj >= m {
						break
					}
					if field[ni][nj] != field[i][j] {
						break
					}
					cnt++
				}
				if cnt >= 5 {
					writer.WriteString("Yes")
					writer.WriteByte('\n')
					return
				}
			}
		}
	}

	writer.WriteString("No")
	writer.WriteByte('\n')
}

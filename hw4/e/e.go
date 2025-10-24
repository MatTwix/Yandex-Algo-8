package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type roadPart struct {
	holes      int64
	importance int64
}

type roadParts []roadPart

func (r roadParts) Len() int {
	return len(r)
}

func (r roadParts) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r roadParts) Less(i, j int) bool {
	return r[i].importance > r[j].importance
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nmkLine, _ := reader.ReadString('\n')
	nmkParts := strings.Fields(nmkLine)
	n, _ := strconv.ParseInt(nmkParts[0], 10, 64)
	m, _ := strconv.ParseInt(nmkParts[1], 10, 64)
	k, _ := strconv.ParseInt(nmkParts[2], 10, 64)

	holesLine, _ := reader.ReadString('\n')
	holesParts := strings.Fields(holesLine)

	holes := make([]int64, n+1)
	for i, v := range holesParts {
		holes[i], _ = strconv.ParseInt(v, 10, 64)
	}

	importances := make([]int64, n+1)

	for range m {
		argLine, _ := reader.ReadString('\n')
		argParts := strings.Fields(argLine)
		l, _ := strconv.Atoi(argParts[0])
		r, _ := strconv.Atoi(argParts[1])

		l--
		r--

		importances[l]++
		importances[r+1]--
	}

	for i := 1; i < int(n); i++ {
		importances[i] = importances[i] + importances[i-1]
	}

	var road roadParts
	var totalDisc int64

	for i := range n {
		road = append(road, roadPart{
			holes:      holes[i],
			importance: importances[i],
		})
		totalDisc += int64(holes[i] * importances[i])
	}

	sort.Sort(road)

	for _, v := range road {
		if k == 0 {
			break
		}

		numToFix := min(k, int64(v.holes))
		k -= numToFix
		totalDisc -= int64(numToFix * v.importance)
	}

	writer.WriteString(strconv.FormatInt(totalDisc, 10))
	writer.WriteByte('\n')
}

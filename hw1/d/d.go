package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	//ввод
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nkLine, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.Fields(nkLine)[1])

	arrLine, _ := reader.ReadString('\n')
	questions := strings.Fields(arrLine)
	// конец ввода

	var topicsMap map[int]int
	topicsMap = make(map[int]int)

	contest := []int{}

	for _, q := range questions {
		qi, _ := strconv.Atoi(q)
		topicsMap[qi]++
	}

	for t := range topicsMap {
		contest = append(contest, t)
		topicsMap[t]--
	}

	for len(contest) < k {
		for t, q := range topicsMap {
			if q == 0 {
				continue
			}
			contest = append(contest, t)
			topicsMap[t]--

			if len(contest) >= k {
				break
			}
		}
	}

	contest = contest[:k]

	for _, q := range contest {
		writer.WriteString(strconv.Itoa(q) + " ")
	}
	writer.WriteByte('\n')
}

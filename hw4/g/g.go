package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type el struct {
	s int
	a int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	sLine, _ := reader.ReadString('\n')
	sParts := strings.Fields(sLine)

	aLine, _ := reader.ReadString('\n')
	aParts := strings.Fields(aLine)

	var elSlice []el

	totalWieght := 0

	for i := range n {
		s, _ := strconv.Atoi(sParts[i])
		a, _ := strconv.Atoi(aParts[i])

		totalWieght += a

		elSlice = append(elSlice, el{s: s, a: a})
	}

	startElSlice := elSlice

	sort.Slice(elSlice, func(i, j int) bool {
		return elSlice[i].s < elSlice[j].s
	})

	curWeight := 0
	e := -1

	for _, v := range elSlice {
		curWeight += v.a
		if curWeight*2 >= totalWieght {
			e = v.s
			break
		}
	}

	minCost := 0

	for _, v := range startElSlice {
		minCost += module(v.s-e) * v.a
	}

	fmt.Fprintf(writer, "%d %d\n", e, minCost)
}

func module(x int) int {
	if x < 0 {
		x = -x
	}

	return x
}

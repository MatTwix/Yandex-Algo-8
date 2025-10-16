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

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	w, _ := strconv.Atoi(parts[1])
	h, _ := strconv.Atoi(parts[2])

	a := make([]int, n)
	b := make([]int, n)
	mnA := 1_000_000_000
	mnB := 1_000_000_000

	for i := range n {
		argLine, _ := reader.ReadString('\n')
		argParts := strings.Fields(argLine)
		ai, _ := strconv.Atoi(argParts[0])
		bi, _ := strconv.Atoi(argParts[1])

		mnA = min(mnA, ai)
		mnB = min(mnB, bi)

		a[i] = ai
		b[i] = bi
	}

	l, r := 0.0, min(float64(w)/float64(mnA), float64(h)/float64(mnB))

	for range 80 {
		m := (l + r) / 2
		if check(m, a, b, w, h) {
			l = m
		} else {
			r = m
		}
	}

	fmt.Printf("%.7f\n", l)
}

func check(k float64, a, b []int, w, h int) bool {
	mxW := float64(w) / k
	totalH := 0.0
	curB := -1
	curSumA := 0.0
	eps := 1e-7

	for i := range len(a) {
		if b[i] != curB {
			if curB != -1 {
				totalH += k * float64(curB)
				if totalH > float64(h)+eps {
					return false
				}
			}
			curB = b[i]
			curSumA = 0
		}

		if float64(a[i]) > mxW+eps {
			return false
		}

		if curSumA+float64(a[i]) <= mxW+eps {
			curSumA += float64(a[i])
		} else {
			totalH += k * float64(curB)
			if totalH > float64(h)+eps {
				return false
			}
			curSumA = float64(a[i])
		}
	}
	if curB != -1 {
		totalH += k * float64(curB)
	}
	return totalH <= float64(h)+eps
}

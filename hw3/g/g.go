package main

import (
	"bufio"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func prefixSums(arr []int64) []int64 {
	prefix := make([]int64, len(arr)+1)
	for i := range arr {
		prefix[i+1] = prefix[i] + arr[i]
	}
	return prefix
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	aLine, _ := reader.ReadString('\n')
	aParts := strings.Fields(aLine)
	aArgs := make([]int64, n)
	for i := range n {
		aArgs[i], _ = strconv.ParseInt(aParts[i], 10, 64)
	}

	mLine, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mLine))

	bLine, _ := reader.ReadString('\n')
	bParts := strings.Fields(bLine)
	bArgs := make([]int64, m)
	for i := range m {
		bArgs[i], _ = strconv.ParseInt(bParts[i], 10, 64)
	}

	a := make([]int64, n)
	copy(a, aArgs)
	slices.Sort(a)

	b := make([]int64, m)
	copy(b, bArgs)
	slices.Sort(b)

	aPrefix := prefixSums(a)
	bPrefix := prefixSums(b)

	var sum1, sum2 int64

	for i := range aArgs {
		currentEl := aArgs[i]
		elInd := sort.Search(len(b), func(i int) bool {
			return b[i] > currentEl
		})

		leftSum := currentEl*int64(elInd) - bPrefix[elInd]
		rightSum := (bPrefix[len(bPrefix)-1] - bPrefix[elInd]) - int64(len(b)-elInd)*currentEl

		sum1 += int64(i+1) * (leftSum + rightSum)
	}

	for i := range bArgs {
		currentEl := bArgs[i]
		elInd := sort.Search(len(a), func(i int) bool {
			return a[i] > currentEl
		})

		leftSum := currentEl*int64(elInd) - aPrefix[elInd]
		rightSum := (aPrefix[len(aPrefix)-1] - aPrefix[elInd]) - int64(len(a)-elInd)*currentEl

		sum2 += int64(i+1) * (leftSum + rightSum)
	}

	writer.WriteString(strconv.Itoa(int(sum1 - sum2)))
	writer.WriteByte('\n')
}

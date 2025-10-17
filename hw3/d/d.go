package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type el struct {
	val int
	idx int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	p, _ := strconv.ParseFloat(parts[1], 64)

	elsLine, _ := reader.ReadString('\n')
	elsParts := strings.Fields(elsLine)

	els := make([]el, n)
	for i, v := range elsParts {
		val, _ := strconv.Atoi(v)
		els[i] = el{val: val, idx: i}
	}

	sort.Slice(els, func(i, j int) bool {
		return els[i].val < els[j].val
	})

	ansI, ansJ := -1, -1
	ansDiff := math.MaxFloat64

	for i, v := range els {
		t := float64(v.val) / p

		j := sort.Search(n, func(k int) bool {
			return float64(els[k].val) >= t
		})

		for _, candidate := range []int{j - 1, j} {
			if candidate < 0 || candidate >= n || candidate == i {
				continue
			}

			r := float64(els[i].val) / float64(els[candidate].val)
			diff := math.Abs(r - p)
			if diff < ansDiff {
				ansDiff = diff
				ansI = els[i].idx
				ansJ = els[candidate].idx
			}
		}
	}

	fmt.Fprintf(writer, "%d %d\n", ansI+1, ansJ+1)
}

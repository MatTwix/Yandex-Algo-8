package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	x int
	y int
}

type vector struct {
	dx int
	dy int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	ndLine, _ := reader.ReadString('\n')
	ndParts := strings.Fields(ndLine)
	n, _ := strconv.Atoi(ndParts[0])
	d, _ := strconv.Atoi(ndParts[1])

	trees := make([]tree, n)
	treeSet := make(map[tree]struct{})

	for i := range n {
		argLine, _ := reader.ReadString('\n')
		argParts := strings.Fields(argLine)
		x, _ := strconv.Atoi(argParts[0])
		y, _ := strconv.Atoi(argParts[1])

		trees[i] = tree{x: x, y: y}
		treeSet[tree{x, y}] = struct{}{}
	}

	vectors := make(map[vector]struct{})

	for dx := 0; dx*dx <= d; dx++ {
		dySq := d - dx*dx
		dy := int(math.Sqrt(float64(dySq)))
		if dy*dy == dySq {
			vectors[vector{dx, dy}] = struct{}{}
			vectors[vector{-dx, dy}] = struct{}{}
			vectors[vector{dx, -dy}] = struct{}{}
			vectors[vector{-dx, -dy}] = struct{}{}
			vectors[vector{dy, dx}] = struct{}{}
			vectors[vector{-dy, dx}] = struct{}{}
			vectors[vector{dy, -dx}] = struct{}{}
			vectors[vector{-dy, -dx}] = struct{}{}
		}
	}

	cnt := 0

	for _, t := range trees {
		for v := range vectors {
			tP := tree{t.x + v.dx, t.y + v.dy}
			if _, exists := treeSet[tP]; exists {
				cnt++
			}
		}
	}

	cnt /= 2

	writer.WriteString(strconv.Itoa(cnt))
	writer.WriteByte('\n')
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int64, n)
	b := make([]int64, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}

	var sumA, sumB int64
	for i := 0; i < n; i++ {
		sumA += a[i]
		sumB += b[i]
	}

	if sumA > sumB {
		fmt.Fprintln(writer, -1)
		return
	}

	left, right := 0, n
	minK := -1

	for left <= right {
		k := (left + right) / 2
		flag := true

		remaining := make([]int64, n)
		copy(remaining, a)
		p := 0

		for i := 0; i < n && flag; i++ {
			for p < n && remaining[p] == 0 {
				p++
			}
			if p < n && p < i-k {
				flag = false
				break
			}
			cap := b[i]
			for cap > 0 {
				for p < n && remaining[p] == 0 {
					p++
				}
				if p >= n || p > i+k {
					break
				}
				take := cap
				if remaining[p] < take {
					take = remaining[p]
				}
				remaining[p] -= take
				cap -= take
				if remaining[p] == 0 {
					p++
				}
			}
		}

		if flag {
			for p < n && remaining[p] == 0 {
				p++
			}
			if p < n {
				flag = false
			}
		}

		if flag {
			minK = k
			right = k - 1
		} else {
			left = k + 1
		}
	}

	writer.WriteString(strconv.Itoa(minK))
	writer.WriteByte('\n')
}

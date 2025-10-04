package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ввод
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])
	// конец ввода

	var penailties map[int]int
	penailties = map[int]int{
		0: -1,
		1: 1,
		2: 0,
		3: 2,
		4: 3,
		5: -1,
		6: 1,
		7: 4,
		8: 2,
		9: 3,
	}

	res := n
	if k != 0 {
		x := res % 10
		res -= x
		if penailties[x] != -1 {
			d := countRemainder(x, penailties[x])

			f := ((k - penailties[x]) / 4) * 20

			i := countRemainder(2, (k-penailties[x])%4) - 2

			// println(penailties[x], ":", d)
			// println(k-penailties[x]-(k-penailties[x])%4, ":", f)
			// println((k-penailties[x])%4, ":", i)
			res += d + f + i
		} else {
			res += x * 2
		}
	}

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}

func countRemainder(start, secs int) int {
	for range secs {
		start += start % 10
	}

	return start
}

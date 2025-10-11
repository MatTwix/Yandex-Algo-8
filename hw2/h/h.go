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

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	res := 1
	if n%4 == 0 {
		res = 2
	}

	writer.WriteString(strconv.Itoa(res))
	writer.WriteByte('\n')
}

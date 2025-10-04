package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ввод
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	arrLine, _ := reader.ReadString('\n')
	parts := strings.Fields(arrLine)
	// конец ввода

	a, _ := strconv.ParseFloat(parts[0], 32)
	b, _ := strconv.ParseFloat(parts[1], 32)
	c, _ := strconv.ParseFloat(parts[2], 32)

	v0, _ := strconv.ParseFloat(parts[3], 32)
	v1, _ := strconv.ParseFloat(parts[4], 32)
	v2, _ := strconv.ParseFloat(parts[5], 32)

	a = min(a, b+c)
	b = min(b, a+c)

	resTime := 0.0

	var1 := a/v0 + a/v1 + b/v0 + b/v1
	var2 := a/v0 + c/v1 + b/v2
	var3 := b/v0 + c/v1 + a/v2

	resTime = min(var1, var2, var3)

	fmt.Fprintf(writer, "%.4f", resTime)
	writer.WriteByte('\n')
}

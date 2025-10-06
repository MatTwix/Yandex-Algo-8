package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	lineStart, _ := reader.ReadString('\n')
	partsStart := strings.Fields(lineStart)
	x, _ := strconv.Atoi(partsStart[0])
	y, _ := strconv.Atoi(partsStart[1])

	lineEnd, _ := reader.ReadString('\n')
	partsEnd := strings.Fields(lineEnd)
	f, _ := strconv.Atoi(partsEnd[0])
	g, _ := strconv.Atoi(partsEnd[1])

	horCRamount := math.Abs(float64(x) - float64(f))
	vertCRamount := math.Abs(float64(y) - float64(g))

	if horCRamount != 0 {
		horCRamount--
	}

	if vertCRamount != 0 {
		vertCRamount--
	}

	res := 3*(horCRamount+vertCRamount) + 1

	if x == f || y == g {
		res--
	}

	fmt.Fprintf(writer, "%.0f", res)
}

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

	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	lDp := 0
	rDp := 1_000_000_000

	for _, c := range s {
		costOfStayLeft := 0
		costOfStayRight := 0

		if rune(c) == 'L' || rune(c) == 'B' {
			costOfStayLeft = 1
		}
		if rune(c) == 'R' || rune(c) == 'B' {
			costOfStayRight = 1
		}

		costOfMoveLeft := 1 + costOfStayLeft
		costOfMoveRight := 1 + costOfStayRight

		newLDp := min(lDp+costOfStayLeft, rDp+costOfMoveLeft)
		newRDp := min(rDp+costOfStayRight, lDp+costOfMoveRight)

		lDp, rDp = newLDp, newRDp
	}

	writer.WriteString(strconv.Itoa(min(rDp, lDp+1)))
	writer.WriteByte('\n')
}

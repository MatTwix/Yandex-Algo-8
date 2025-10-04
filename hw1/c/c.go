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
	// ввод

	var runeMap map[rune]int
	runeMap = make(map[rune]int)

	for _, r := range s {
		runeMap[r]++
	}

	pairs := len(s) * (len(s) - 1) / 2
	samePairs := 0
	for _, cnt := range runeMap {
		samePairs += cnt * (cnt - 1) / 2
	}

	// вывод
	writer.WriteString(strconv.Itoa(pairs - samePairs + 1))
	writer.WriteByte('\n')
}

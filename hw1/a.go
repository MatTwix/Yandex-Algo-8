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

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	arrLine, _ := reader.ReadString('\n')
	parts := strings.Fields(arrLine)
	// конец ввода

	happyness := 0
	mnV, mxM := 1000, 0

	for i := range n {
		mush, _ := strconv.Atoi(parts[i])
		if i%2 == 0 {
			happyness += mush
			mnV = min(mnV, mush)
		} else {
			happyness -= mush
			mxM = max(mxM, mush)
		}
	}

	mx := max(mxM-mnV, 0)

	happyness += mx * 2

	// вывод

	// for _, mush := range m {
	// 	fmt.Println(mush)
	// }

	// fmt.Println("----")

	// for _, mush := range v {
	// 	fmt.Println(mush)
	// }

	writer.WriteString(strconv.Itoa(happyness))
	writer.WriteByte('\n')
	// конец вывода
}

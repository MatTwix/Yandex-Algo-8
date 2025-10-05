package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	sLine, _ := reader.ReadString('\n')
	s := strings.TrimSpace(sLine)

	pieceLength := n / m

	reps := make(map[string][]int)

	for i := range m {
		pieceLine, _ := reader.ReadString('\n')
		piece := strings.TrimSpace(pieceLine)
		reps[piece] = append(reps[piece], i+1)
	}

	res := make([]int, 0, m)

	for i := range m {
		currentPiece := s[i*pieceLength : (i+1)*pieceLength]
		indices := reps[currentPiece]

		validIndex := -1
		for ind, x := range indices {
			if x >= 0 {
				validIndex = x
				indices[ind] = -1
				break
			}
		}

		res = append(res, validIndex)
		reps[currentPiece] = indices
	}

	for i, val := range res {
		if i > 0 {
			writer.WriteByte(' ')
		}
		writer.WriteString(strconv.Itoa(val))
	}
	writer.WriteByte('\n')
}

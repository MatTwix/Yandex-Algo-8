package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type store struct {
	retPrice int
	wsPrice  int
	wsAmount int
	inStock  int
}

type dpVal struct {
	val        int
	currAmount int
	prevAmount int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	n, _ := strconv.Atoi(parts[0])
	l, _ := strconv.Atoi(parts[1])
	// конец ввода

	stores := make([]store, n)
	for i := range n {
		sLine, _ := reader.ReadString('\n')
		sParts := strings.Fields(sLine)

		stores[i].retPrice, _ = strconv.Atoi(sParts[0])
		stores[i].wsAmount, _ = strconv.Atoi(sParts[1])
		stores[i].wsPrice, _ = strconv.Atoi(sParts[2])
		stores[i].inStock, _ = strconv.Atoi(sParts[3])
	}

	dp := make([][]dpVal, n)
	for i := range n {
		dp[i] = make([]dpVal, 101)
	}

	for i := range dp[0] {
		dp[0][i].val = countPrice(
			stores[0].retPrice,
			stores[0].wsAmount,
			stores[0].wsPrice,
			stores[0].inStock,
			i,
		)
		dp[0][i].currAmount = i
	}

	for i := 1; i < n; i++ {
		for j := range 101 {
			soloPrice := countPrice(
				stores[i].retPrice,
				stores[i].wsAmount,
				stores[i].wsPrice,
				stores[i].inStock,
				j,
			)

			mn := soloPrice
			currAmount := j
			prevAmount := 0

			dp[i][j].val = mn
			dp[i][j].currAmount = j

			for k := range j + 1 {
				crPrice := countPrice(
					stores[i].retPrice,
					stores[i].wsAmount,
					stores[i].wsPrice,
					stores[i].inStock,
					j-k,
				)

				if mn > dp[i-1][k].val+crPrice {
					mn = dp[i-1][k].val + crPrice
					currAmount = j - k
					prevAmount = k
				}
			}

			dp[i][j].val = mn
			dp[i][j].currAmount = currAmount
			dp[i][j].prevAmount = prevAmount
		}
	}

	ans := dp[n-1][l].val

	for i := l; i < 101; i++ {
		if dp[n-1][i].val < ans {
			ans = dp[n-1][i].val
			l = i
		}
	}

	if ans == 1000000000 {
		ans = -1
		writer.WriteString(strconv.Itoa(ans))
		writer.WriteByte('\n')

		return
	}

	ansRec := make([]string, n)

	currAmount := l
	for i := n - 1; i >= 0; i-- {
		ansRec[i] = strconv.Itoa(dp[i][currAmount].currAmount)
		currAmount = dp[i][currAmount].prevAmount
	}

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
	writer.WriteString(strings.Join(ansRec, " "))
	writer.WriteByte('\n')
}

func countPrice(retPrice, wsAmount, wsPrice, inStock, amount int) int {
	if amount > inStock {
		return 1_000_000_000
	}

	if amount < wsAmount {
		return amount * retPrice
	}

	return amount * wsPrice

}

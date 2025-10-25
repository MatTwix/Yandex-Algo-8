package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	Enter = 1
	Leave = -1
)

type train struct {
	a int
	b int
	v int
}

type event struct {
	Time float64
	Type int
}

type eventSlice []event

func (s eventSlice) Len() int {
	return len(s)
}

func (s eventSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s eventSlice) Less(i, j int) bool {
	if s[i].Time != s[j].Time {
		return s[i].Time < s[j].Time
	}

	return s[i].Type < s[j].Type
}

type car struct {
	Time int
	Idx  int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nmxLine, _ := reader.ReadString('\n')
	nmxParts := strings.Fields(nmxLine)
	n, _ := strconv.Atoi(nmxParts[0])
	m, _ := strconv.Atoi(nmxParts[1])
	x, _ := strconv.Atoi(nmxParts[2])

	var trains []train

	for range n {
		argLine, _ := reader.ReadString('\n')
		argParts := strings.Fields(argLine)
		a, _ := strconv.Atoi(argParts[0])
		b, _ := strconv.Atoi(argParts[1])
		v, _ := strconv.Atoi(argParts[2])

		trains = append(trains, train{a: a, b: b, v: v})
	}

	cars := make([]car, m)

	carsLine, _ := reader.ReadString('\n')
	carsParts := strings.Fields(carsLine)

	for i := range m {
		cars[i].Time, _ = strconv.Atoi(carsParts[i])
		cars[i].Idx = i
	}

	var events eventSlice

	for _, v := range trains {
		var directedV float64
		if v.a < v.b {
			directedV = float64(v.v)
		} else {
			directedV = -float64(v.v)
		}

		t1 := float64(x-v.a) / directedV
		t2 := float64(x-v.b) / directedV

		tEnter := minF(t1, t2)
		tLeave := maxF(t1, t2)

		if tLeave >= 0 {
			events = append(events, event{Time: maxF(0.0, tEnter), Type: Enter})
			events = append(events, event{Time: tLeave, Type: Leave})
		}
	}

	sort.Sort(events)

	var intervals [][]float64
	balance := 0
	curTime := 0.0

	for _, v := range events {
		if v.Time > curTime && balance == 0 {
			intervals = append(intervals, []float64{curTime, v.Time})
		}
		balance += v.Type
		curTime = v.Time
	}

	intervals = append(intervals, []float64{curTime, math.MaxFloat64})

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Time < cars[j].Time
	})
	ans := make([]float64, m)

	freeInterval := 0

	for _, v := range cars {
		for freeInterval < len(intervals) && intervals[freeInterval][1] <= float64(v.Time) {
			freeInterval++
		}

		if freeInterval < len(intervals) {
			ans[v.Idx] = maxF(float64(v.Time), intervals[freeInterval][0])
		}
	}

	for _, v := range ans {
		fmt.Fprintf(writer, "%.6f\n", v)
	}
}

func maxF(i, j float64) float64 {
	if i > j {
		return i
	}

	return j
}

func minF(i, j float64) float64 {
	if i < j {
		return i
	}

	return j
}

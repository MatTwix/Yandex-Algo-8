package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	Arrival   = 1
	Departure = 2
)

type event struct {
	Time int
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

func parseTime(timeStr string) int {
	parts := strings.Split(timeStr, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])

	return hours*60 + minutes
}

func parseLine(line string) (start, end int) {
	parts := strings.Split(line, "-")

	start = parseTime(parts[0])
	end = parseTime(parts[1])

	return start, end
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))

	var firstOfficeEvents eventSlice
	var secondOffiseEvents eventSlice

	for range n {
		argLine, _ := reader.ReadString('\n')
		start, end := parseLine(strings.TrimSpace(argLine))

		firstOfficeEvents = append(firstOfficeEvents, event{Time: start, Type: Departure})
		secondOffiseEvents = append(secondOffiseEvents, event{Time: end, Type: Arrival})
	}

	mLine, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mLine))

	for range m {
		argLine, _ := reader.ReadString('\n')
		start, end := parseLine(strings.TrimSpace(argLine))

		secondOffiseEvents = append(secondOffiseEvents, event{Time: start, Type: Departure})
		firstOfficeEvents = append(firstOfficeEvents, event{Time: end, Type: Arrival})
	}

	sort.Sort(firstOfficeEvents)
	sort.Sort(secondOffiseEvents)

	ans := processEvents(firstOfficeEvents) + processEvents(secondOffiseEvents)

	writer.WriteString(strconv.Itoa(ans))
	writer.WriteByte('\n')
}

func processEvents(officeEvents eventSlice) int {
	bussesNeeded := 0
	bussesAvailable := 0

	for i := range officeEvents {
		if officeEvents[i].Type == Departure {
			if bussesAvailable > 0 {
				bussesAvailable--
			} else if bussesAvailable == 0 {
				bussesNeeded++
			}
		} else {
			bussesAvailable++
		}
	}

	return bussesNeeded
}

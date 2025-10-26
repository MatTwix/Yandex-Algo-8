package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	Finish    = 1
	Collision = 2
	Wall      = 3
	epsilon   = 1e-9
)

type car struct {
	x  int
	y  int
	vX int
	vY int
}

type event struct {
	Time float64
	Type int
	Idx1 int
	Idx2 int
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	nlwLine, _ := reader.ReadString('\n')
	nlwParts := strings.Fields(nlwLine)
	n, _ := strconv.Atoi(nlwParts[0])
	l, _ := strconv.Atoi(nlwParts[1])
	w, _ := strconv.Atoi(nlwParts[2])

	cars := make([]car, n)

	for i := range n {
		argLine, _ := reader.ReadString('\n')
		argParts := strings.Fields(argLine)
		x, _ := strconv.Atoi(argParts[0])
		y, _ := strconv.Atoi(argParts[1])
		vX, _ := strconv.Atoi(argParts[2])
		vY, _ := strconv.Atoi(argParts[3])

		cars[i] = car{
			x:  x,
			y:  y,
			vX: vX,
			vY: vY,
		}
	}

	var events []event

	for i, v := range cars {
		tFinish := math.MaxFloat64

		if v.vX != 0 {
			tFinish = float64(l-v.x) / float64(v.vX)
		}

		var tWall float64

		switch {
		case v.vY > 0:
			tWall = float64(w-v.y) / float64(v.vY)
		case v.vY < 0:
			tWall = float64(-v.y) / float64(v.vY)
		default:
			tWall = math.MaxFloat64
		}

		for j := i + 1; j < n; j++ {
			dx := cars[j].x - v.x
			dvX := v.vX - cars[j].vX

			dy := cars[j].y - v.y
			dvY := v.vY - cars[j].vY

			t := math.MaxFloat64

			switch {
			case dvX != 0 && dvY != 0:
				tX := float64(dx) / float64(dvX)
				tY := float64(dy) / float64(dvY)

				if areEqual(tX, tY) {
					t = tX
				}
			case dvX == 0:
				if dx == 0 && dvY != 0 {
					tY := float64(dy) / float64(dvY)
					if tY > 0 {
						t = tY
					}
				}
			case dvY == 0:
				if dy == 0 && dvX != 0 {
					tX := float64(dx) / float64(dvX)
					if tX > 0 {
						t = tX
					}
				}
			}

			if t > epsilon && t != math.MaxFloat64 {
				events = append(events, event{Time: t, Type: Collision, Idx1: i, Idx2: j})
			}
		}

		if tFinish > epsilon && tFinish != math.MaxFloat64 {
			events = append(events, event{Time: tFinish, Type: Finish, Idx1: i})
		}

		if tWall > epsilon && tWall != math.MaxFloat64 {
			events = append(events, event{Time: tWall, Type: Wall, Idx1: i})
		}
	}

	sort.Slice(events, func(i, j int) bool {
		if !areEqual(events[i].Time, events[j].Time) {
			return events[i].Time < events[j].Time
		}

		return events[i].Type > events[j].Type
	})

	isEliminated := make([]bool, n)
	finishes := make([]float64, n)

	for i := range finishes {
		finishes[i] = math.MaxFloat64
	}

	for i := 0; i < len(events); {
		j := i
		for j+1 < len(events) && areEqual(events[i].Time, events[j+1].Time) {
			j++
		}

		toBeEliminated := make(map[int]bool)

		for k := i; k <= j; k++ {
			event := events[k]
			switch event.Type {
			case Wall:
				toBeEliminated[event.Idx1] = true
			case Collision:
				if !isEliminated[event.Idx1] && !isEliminated[event.Idx2] {
					toBeEliminated[event.Idx1] = true
					toBeEliminated[event.Idx2] = true
				}
			}
		}

		for k := i; k <= j; k++ {
			event := events[k]
			if event.Type == Finish {
				if !isEliminated[event.Idx1] && !toBeEliminated[event.Idx1] {
					finishes[event.Idx1] = event.Time
				}
			}
		}

		for idx := range toBeEliminated {
			isEliminated[idx] = true
		}

		i = j + 1
	}

	minFinishTime := math.MaxFloat64

	for _, v := range finishes {
		minFinishTime = minF(minFinishTime, v)
	}

	if minFinishTime == math.MaxFloat64 {
		writer.WriteString("0")
		writer.WriteByte('\n')
		return
	}

	var ans []string

	for i, v := range finishes {
		if areEqual(v, minFinishTime) {
			ans = append(ans, strconv.Itoa(i+1))
		}
	}

	writer.WriteString(strconv.Itoa(len(ans)))
	writer.WriteByte('\n')
	writer.WriteString(strings.Join(ans, " "))
	writer.WriteByte('\n')
}

func areEqual(i, j float64) bool {
	return math.Abs(i-j) < epsilon
}

func minF(i, j float64) float64 {
	if i < j {
		return i
	}

	return j
}

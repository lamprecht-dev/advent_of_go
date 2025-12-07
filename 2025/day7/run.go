package main

import (
	"aog/internal/aogutils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	data := aogutils.GetInput()
	result, result2 := solve(data)
	fmt.Println(result, result2)
}

type beam struct {
	Pos   int
	Times int
}

func (b *beam) In(beams *[]beam) int {
	for i, b2 := range *beams {
		if b.Pos == b2.Pos {
			return i
		}
	}
	return -1
}

func (b *beam) Integrate(beams *[]beam) []beam {
	if i := b.In(beams); i != -1 {
		(*beams)[i].Times += b.Times
	} else {
		return append((*beams), *b)
	}
	return *beams
}

func parse_data(data string) (currentBeams []beam, splitters [][]int) {
	for _, row := range strings.Split(data, "\n") {
		rowSplitters := make([]int, 0)
		for x := 0; x < len(row); x++ {
			cell := string(row[x])
			switch cell {
			case "S":
				currentBeams = append(currentBeams, beam{Pos: x, Times: 1})
			case "^":
				rowSplitters = append(rowSplitters, x)
			}
		}
		if len(rowSplitters) > 0 {
			splitters = append(splitters, rowSplitters)
		}
	}
	return
}

func solve(data string) (splits int, timelines int) {
	currentBeams, splitters := parse_data(data)
	timelines = 1
	splits = 0

	for _, s := range splitters {
		newBeams := make([]beam, 0)
		for _, b := range currentBeams {
			if slices.Contains(s, b.Pos) {
				timelines += b.Times
				splits += 1
				l := beam{Pos: b.Pos - 1, Times: b.Times}
				r := beam{Pos: b.Pos + 1, Times: b.Times}
				newBeams = l.Integrate(&newBeams)
				newBeams = r.Integrate(&newBeams)
			} else {
				newBeams = b.Integrate(&newBeams)
			}
		}
		currentBeams = newBeams
		// fmt.Println(s, currentBeams, splits, timelines)
	}

	return
}

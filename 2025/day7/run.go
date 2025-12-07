package main

import (
	"aog/internal/aogutils"
	"fmt"
	"strings"
)

func main() {
	data := aogutils.GetInput()
	result, result2 := solve(data)
	fmt.Println(result, result2)
}

func parseData(data string) (currentBeams map[int]int, splitters []map[int]struct{}) {
	currentBeams = make(map[int]int)
	for _, row := range strings.Split(data, "\n") {
		rowSplitters := make(map[int]struct{})
		for x := 0; x < len(row); x++ {
			switch row[x] {
			case 'S':
				currentBeams[x] = 1
			case '^':
				rowSplitters[x] = struct{}{}
			}
		}
		if len(rowSplitters) > 0 {
			splitters = append(splitters, rowSplitters)
		}
	}
	return
}

func solve(data string) (splits int, timelines int) {
	currentBeams, splitters := parseData(data)
	timelines = 1

	for _, s := range splitters {
		newBeams := make(map[int]int)
		for pos, times := range currentBeams {
			if _, ok := s[pos]; ok {
				timelines += times
				splits += 1
				newBeams[pos-1] += times
				newBeams[pos+1] += times
			} else {
				newBeams[pos] += times
			}
		}
		currentBeams = newBeams
	}

	return
}

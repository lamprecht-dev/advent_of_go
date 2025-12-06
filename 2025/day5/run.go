package main

import (
	"aog/internal/aogutils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := aogutils.GetInput()
	result, result2 := solve(data)
	fmt.Println(result, result2)
}

type myRange struct {
	min int
	max int
}

func (r *myRange) contains(v int) bool {
	return v <= r.max && v >= r.min
}

func (r *myRange) overlaps(r2 *myRange) bool {
	return r.max >= r2.min && r.min <= r2.max
}

func (r *myRange) count() int {
	return r.max - r.min + 1
}

func (r *myRange) extend(r2 *myRange) {
	r.min = min(r.min, r2.min)
	r.max = max(r.max, r2.max)
}

func solve(data string) (s1 int, s2 int) {
	ranges := make([]myRange, 0)
	findingRanges := true
	for _, row := range strings.Split(data, "\n") {
		if findingRanges {
			if row == "" {
				findingRanges = false
				continue
			}
			l, r, _ := strings.Cut(row, "-")
			min, _ := strconv.Atoi(l)
			max, _ := strconv.Atoi(r)
			ranges = append(ranges, myRange{min, max})
		} else {
			v, _ := strconv.Atoi(row)
			for _, r := range ranges {
				if r.contains(v) {
					s1 += 1
					break
				}
			}
		}
	}

	newRanges := combineRanges(ranges)
	for len(newRanges) < len(ranges) {
		ranges = newRanges
		newRanges = combineRanges(ranges)
	}

	for _, r := range ranges {
		s2 += r.count()
	}

	return
}

func combineRanges(ranges []myRange) (newRanges []myRange) {
	newRanges = make([]myRange, 0)
	for _, r := range ranges {
		added := false
		for i, nr := range newRanges {
			if nr.overlaps(&r) {
				newRanges[i].extend(&r)
				added = true
				break
			}
		}
		if added {
			continue
		} else {
			newRanges = append(newRanges, r)
		}
	}
	return
}

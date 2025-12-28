package main

import (
	"aog/internal/aogutils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := aogutils.GetInput()
	result := solve1(data)
	result2 := solve2(data)
	fmt.Println(result, result2)
}

type Pos struct {
	X int
	Y int
}

// func parse(data string) map[int][]int {
func parse(data string) []Pos {
	// pos := make(map[int][]int)
	pos := make([]Pos, 0)
	for _, d := range strings.Split(data, "\n") {
		dsplit := strings.Split(d, ",")
		x, _ := strconv.Atoi(dsplit[0])
		y, _ := strconv.Atoi(dsplit[1])
		pos = append(pos, Pos{X: x, Y: y})
		// pos[y] = append(pos[y], x)
	}
	return pos
}

func solve1(data string) (maxArea int) {
	pos := parse(data)
	for i, p1 := range pos {
		for j, p2 := range pos {
			if j <= i {
				continue
			}
			area := (math.Abs(float64(p2.X-p1.X)) + 1) * (math.Abs(float64(p2.Y-p1.Y)) + 1)
			maxArea = max(maxArea, int(area))
		}
	}
	return
}

func sortedKeys(m map[int]struct{}) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func parsePos(pos []Pos) ([]int, []int) {
	xSet := make(map[int]struct{})
	ySet := make(map[int]struct{})

	for _, p := range pos {
		xSet[p.X] = struct{}{}
		ySet[p.Y] = struct{}{}
	}

	return sortedKeys(xSet), sortedKeys(ySet)
}

type Rect struct {
	pos1 Pos // min
	pos2 Pos // max
}

func isInside(rect Rect, vertWalls map[Rect]struct{}) bool {
	mX := (rect.pos1.X + rect.pos2.X) / 2
	mY := (rect.pos1.Y + rect.pos2.Y) / 2

	count := 0
	for vw := range vertWalls {
		minW := min(vw.pos1.Y, vw.pos2.Y)
		maxW := max(vw.pos1.Y, vw.pos2.Y)
		if vw.pos1.X > mX && minW < mY && maxW > mY {
			count += 1
		}
	}

	return count%2 == 1
}

func overlapps(r1 Rect, r2 Rect) bool {
	minR1X := min(r1.pos1.X, r1.pos2.X)
	maxR1X := max(r1.pos1.X, r1.pos2.X)

	minR2X := min(r2.pos1.X, r2.pos2.X)
	maxR2X := max(r2.pos1.X, r2.pos2.X)

	minR1Y := min(r1.pos1.Y, r1.pos2.Y)
	maxR1Y := max(r1.pos1.Y, r1.pos2.Y)

	minR2Y := min(r2.pos1.Y, r2.pos2.Y)
	maxR2Y := max(r2.pos1.Y, r2.pos2.Y)
	return minR1X < maxR2X && minR2X < maxR1X && minR1Y < maxR2Y && minR2Y < maxR1Y
}

func solve2(data string) (maxArea int) {
	pos := parse(data)
	xList, yList := parsePos(pos)
	var rects map[Rect]bool = make(map[Rect]bool)

	vertWalls := make(map[Rect]struct{})
	for i := 0; i < len(pos); i++ {
		if i == len(pos)-1 {
			vertWalls[Rect{pos1: pos[i], pos2: pos[0]}] = struct{}{}
			break
		}
		if pos[i].X == pos[i+1].X {
			vertWalls[Rect{pos1: pos[i], pos2: pos[i+1]}] = struct{}{}
		}
	}

	for xi := 0; xi < len(xList)-1; xi++ {
		for yi := 0; yi < len(yList)-1; yi++ {
			rect := Rect{pos1: Pos{X: xList[xi], Y: yList[yi]}, pos2: Pos{X: xList[xi+1], Y: yList[yi+1]}}
			rects[rect] = isInside(rect, vertWalls)
		}
	}

	// Now finally we do the same as solve1 but check if all rects that are included in here are inside the poligon
	for i, p1 := range pos {
	innerLoop:
		for j, p2 := range pos {
			if j <= i {
				continue innerLoop
			}

			rect := Rect{pos1: p1, pos2: p2}

			for r, inside := range rects {
				if overlapps(r, rect) && !inside {
					continue innerLoop
				}
			}

			area := (math.Abs(float64(p2.X-p1.X)) + 1) * (math.Abs(float64(p2.Y-p1.Y)) + 1)
			maxArea = max(maxArea, int(area))
		}
	}
	return
}

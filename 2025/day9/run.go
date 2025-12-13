package main

import (
	"aog/internal/aogutils"
	"fmt"
	"math"
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

func solve2(data string) (maxArea int) {
	pos := parse(data)
	horLines := make(map[int][]Pos) // Abusing Pos for the span of the line
	verLines := make(map[int][]Pos) // Abusing Pos for the span of the line
	globalMinY, globalMaxY, globalMinX, globalMaxX := math.MaxInt32, 0, math.MaxInt32, 0
	for i := 0; i < len(pos); i++ {
		p1 := Pos{}
		p2 := Pos{}
		if i == len(pos)-1 {
			p1 = pos[i]
			p2 = pos[0]
		} else {
			p1 = pos[i]
			p2 = pos[i+1]
		}
		if p1 == p2 {
			continue // sanity check
		}
		if p1.X == p2.X {
			verLines[p1.X] = append(verLines[p1.X], Pos{min(p1.Y, p2.Y), max(p1.Y, p2.Y)})
			globalMaxX = max(globalMaxX, p1.X)
			globalMinX = min(globalMinX, p1.X)
		} else {
			horLines[p1.Y] = append(horLines[p1.Y], Pos{min(p1.X, p2.X), max(p1.X, p2.X)})
			globalMaxY = max(globalMaxY, p1.Y)
			globalMinY = min(globalMinY, p1.Y)
		}
	}

	// Same as solve1 BUT now we check:
	// Is no line crossing through this area?
	// AND an inside point looks at 4 walls
	for i, p1 := range pos {
	p2Break:
		for j, p2 := range pos {
			if j <= i {
				continue
			}
			minY, maxY, minX, maxX := min(p1.Y, p2.Y), max(p1.Y, p2.Y), min(p1.X, p2.X), max(p1.X, p2.X)

			for y := minY + 1; y < maxY; y++ {
				if lines, ok := horLines[y]; ok {
					for _, line := range lines {
						minX2, maxX2 := line.X, line.Y
						if minX < maxX2 && minX2 < maxX {
							break p2Break
						}
					}
				}
			}

			for x := minX + 1; x < maxX; x++ {
				if lines, ok := verLines[x]; ok {
					for _, line := range lines {
						minY2, maxY2 := line.X, line.Y
						if minY < maxY2 && minY2 < maxY {
							break p2Break
						}
					}
				}
			}
			innerW := maxX - minX - 1
			innerH := maxY - minY - 1

			// If no interior exists, skip the 4-wall test;
			// line-crossing test is enough.
			if innerW < 1 || innerH < 1 {
				area := (maxX - minX + 1) * (maxY - minY + 1)
				maxArea = max(maxArea, area)
				continue
			}

			// Final condition, an inside point must see 4 walls!
			origX := minX + 1
			origY := minY + 1
			foundEdge := 0

			ipy := origY
		wall1:
			for ipx := origX; ipx <= globalMaxX; ipx++ {
				if lines, ok := verLines[ipx]; ok {
					for _, line := range lines {
						if line.X <= ipy && line.Y >= ipy {
							foundEdge++
							break wall1
						}
					}
				}
			}

			if foundEdge != 1 {
				break p2Break
			}

			ipy = origY
		wall2:
			for ipx := origX; ipx >= globalMinX; ipx-- {
				if lines, ok := verLines[ipx]; ok {
					for _, line := range lines {
						if line.X <= ipy && line.Y >= ipy {
							foundEdge++
							break wall2
						}
					}
				}
			}

			if foundEdge != 2 {
				break p2Break
			}

			ipx := origX
		wall3:
			for ipy := origY; ipy <= globalMaxY; ipy++ {
				if lines, ok := horLines[ipy]; ok {
					for _, line := range lines {
						if line.X <= ipx && line.Y >= ipx {
							foundEdge++
							break wall3
						}
					}
				}
			}

			if foundEdge != 3 {
				break p2Break
			}

			ipx = origX
		wall4:
			for ipy := origY; ipy >= globalMinY; ipy-- {
				if lines, ok := horLines[ipy]; ok {
					for _, line := range lines {
						if line.X <= ipx && line.Y >= ipx {
							foundEdge++
							break wall4
						}
					}
				}
			}

			if foundEdge != 4 {
				break p2Break
			}

			area := (maxX - minX + 1) * (maxY - minY + 1)
			maxArea = max(maxArea, int(area))
		}
	}
	return
}

// 114121316 too low

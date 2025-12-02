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
	fmt.Println(solve(data))
}

func solve(data string) (int, int) {
	pos := 50
	count0a := 0
	count0b := 0.0
	for _, line := range strings.Split(data, "\n") {
		oldPos := pos
		n, _ := strconv.Atoi(line[1:])          // Really bad practise to ignore errors, but hey, its aoc!
		count0b += math.Floor(float64(n) / 100) // guaranteed 0s!
		if line[:1] == "R" {
			pos = (pos + n) % 100
			if pos > 99 {
				pos -= 100
			}
			if pos < oldPos {
				count0b++
			}
		} else {
			pos = (pos - n) % 100
			if pos < 0 {
				pos += 100
			}
			if oldPos != 0 && (pos == 0 || pos > oldPos) {
				count0b++
			}
		}
		if pos == 0 {
			count0a++
		}
	}
	return count0a, int(count0b)
}

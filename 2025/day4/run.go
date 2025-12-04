package main

import (
	"aog/internal/aogutils"
	"fmt"
	"strings"
)

func main() {
	data := aogutils.GetInput()
	result := solve1(data)
	result2 := solve2(data)
	fmt.Println(result, result2)
}

func removeRolls(data string, continues bool) (rolls int) {
	rows := strings.Split(data, "\n")
	changes := true
	for changes {
		changes = false
		for y, row := range rows {
		col:
			for x, cell := range strings.Split(row, "") {
				if cell == "." {
					continue
				}
				cellCount := 0
				for _, offset := range aogutils.Dirs {
					ny := y + offset[0]
					nx := x + offset[1]
					if ny < 0 || ny >= len(rows) || nx < 0 || nx >= len(rows[0]) {
						continue
					}
					if rows[ny][nx] == '@' {
						cellCount += 1
					}
					if cellCount >= 4 {
						continue col // Gotta love that go uses gotos!
					}
				}
				if cellCount <= 3 {
					rolls += 1
					if continues {
						old := []byte(rows[y])
						old[x] = '.'
						rows[y] = string(old)
						changes = true
					}
				}
			}
		}
	}
	return
}

func solve1(data string) (rolls int) {
	return removeRolls(data, false)
}

func solve2(data string) (rolls int) {
	return removeRolls(data, true)
}

package main

import (
	"aog/internal/aogutils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := aogutils.GetInput()
	result := solve1(data)
	result2 := solve2(data)
	fmt.Println(result, result2)
}

func parse_row(data string) (conf []int, buttons []map[int]bool, jolt []int) {
	objs := strings.Split(data, " ")
	buttons = make([]map[int]bool, 0)
	for _, obj := range objs {
		switch obj[0] {
		case '[':
			binary_data := obj[1 : len(obj)-1]
			for _, bd := range binary_data {
				if bd == '#' {
					conf = append(conf, 1)
				} else {
					conf = append(conf, 0)
				}
			}
		case '(':
			binary_data := strings.Split(obj[1:len(obj)-1], ",")
			button := make(map[int]bool)
			for _, bd := range binary_data {
				i, _ := strconv.Atoi(bd)
				button[i] = true
			}
			buttons = append(buttons, button)
		case '{':
			jolts := strings.Split(obj[1:len(obj)-1], ",")
			jolt = make([]int, len(jolts))
			for i, j := range jolts {
				jolt[i], _ = strconv.Atoi(j)
			}
		}
	}

	return
}

func solve1(data string) (sum int) {
	// rows := strings.Split(data, "\n")
	// for i := 0; i < len(rows); i++ {
	// 	c, b, j := parse_row(rows[i])
	// 	fmt.Println(c, b, j)
	// }
	return
}

type State struct {
	Goal  []int
	Count int
}

func solve2(data string) (sum int) {
	rows := strings.Split(data, "\n")
rowLoop:
	for i := 0; i < len(rows); i++ {
		_, buttons, j := parse_row(rows[i])

		next := []State{{Goal: j, Count: 0}}
		seen := make(map[string]bool, 0)
		for len(next) > 0 {
			cur := next[0]
			next = next[1:]
			curS := fmt.Sprint(cur.Goal)

			if isCorrect(cur.Goal) {
				sum += cur.Count
				continue rowLoop
			}

			if _, ok := seen[curS]; ok {
				continue
			}
			seen[curS] = true

		butonLoop:
			for _, button := range buttons {
				nextGoal := make([]int, len(cur.Goal))
				copy(nextGoal, cur.Goal)
				nextState := State{Goal: nextGoal, Count: cur.Count + 1}
				for b := range button {
					nextState.Goal[b]--
					if nextState.Goal[b] < 0 {
						continue butonLoop
					}
				}
				next = append(next, nextState)
			}
		}
	}
	return
}

func isCorrect(jolts []int) bool {
	for _, j := range jolts {
		if j != 0 {
			return false
		}
	}
	return true
}

// func solve3(data string) (sum int) {
// 	rows := strings.Split(data, "\n")
// 	for i := 0; i < len(rows); i++ {
// 		_, b, j := parse_row(rows[i])
// 		cache := make(map[string]int)
// 		sum += solve_jolts(j, b, cache)
// 	}
// 	return
// }

// func solve_jolts(target []int, buttons []map[int]bool, cache map[string]int) int {
// 	fmt.Println(target)
// 	if v, ok := cache[fmt.Sprint(target)]; ok {
// 		return v
// 	}
// 	last := true
// 	for t := range target {
// 		if t != 0 {
// 			last = false
// 			break
// 		}
// 	}
// 	if last {
// 		return 0
// 	}

// 	best := math.MaxInt
// buttonLoop:
// 	for _, b := range buttons {
// 		new_target := make([]int, len(target))
// 		for i := 0; i < len(target); i++ {
// 			target_val := target[i]
// 			if _, ok := b[i]; ok {
// 				target_val--
// 				if target_val < 0 {
// 					continue buttonLoop
// 				}
// 			}
// 			new_target[i] = target_val
// 		}
// 		cur := solve_jolts(new_target, buttons, cache)
// 		best = min(cur, best)
// 	}
// 	if best < 0 {
// 		best = math.MaxInt
// 	}
// 	cache[fmt.Sprint(target)] = best
// 	return best
// }

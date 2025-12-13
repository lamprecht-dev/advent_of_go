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

type bm uint64 // bitmask

func parse_row(data string) (conf bm, buttons []bm, jolt []int) {
	objs := strings.Split(data, " ")
	buttons = make([]bm, len(objs)-2)
	bi := 0
	for _, obj := range objs {
		switch obj[0] {
		case '[':
			binary_data := obj[1 : len(obj)-1]
			for i, bd := range binary_data {
				if bd == '#' {
					conf |= 1 << i // set mask at pos i
				}
			}
		case '(':
			binary_data := strings.Split(obj[1:len(obj)-1], ",")
			conf := 0
			for _, bd := range binary_data {
				i, _ := strconv.Atoi(bd)
				conf |= 1 << i
			}
			buttons[bi] = bm(conf)
			bi++
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
	for _, row := range strings.Split(data, "\n") {
		conf, buttons, _ := parse_row(row)
		sum += find_combination(conf, buttons, 0, make(map[cacheKey]int))
	}
	return
}

type cacheKey struct {
	target  bm
	pressed bm
}

func find_combination(target bm, buttons []bm, pressed bm, cache map[cacheKey]int) (best int) {
	if v, ok := cache[cacheKey{target: target, pressed: pressed}]; ok {
		return v
	}
	if target == 0 {
		return 0
	}
	best = math.MaxInt
	for i, b := range buttons {
		if pressed&(1<<i) != 0 {
			continue
		}
		new_pressed := pressed
		new_pressed |= 1 << i

		cur := find_combination(target^b, buttons, new_pressed, cache)
		if cur != math.MaxInt {
			best = min(best, cur+1)
		}
	}
	cache[cacheKey{target: target, pressed: pressed}] = best
	return best
}

func solve2(data string) (sum int) {
	return
}

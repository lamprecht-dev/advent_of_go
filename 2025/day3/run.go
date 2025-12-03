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

// My initial solution for part 1, below replaced with the refactor from part 2
// func solve1(data string) (sum int) {
// 	for _, bank := range strings.Split(data, "\n") {
// 		largest := 0
// 		for l := 0; l < len(bank); l++ {
// 			lDig, _ := strconv.Atoi(bank[l : l+1])
// 			if lDig*10+9 <= largest {
// 				continue
// 			}
// 			for r := l + 1; r < len(bank); r++ {
// 				rDig, _ := strconv.Atoi(bank[r : r+1])
// 				largest = max(largest, lDig*10+rDig)
// 			}
// 		}
// 		sum += largest
// 	}
// 	return sum
// }

func solve1(data string) (sum int) {
	for _, bank := range strings.Split(data, "\n") {
		sum += findSubJolt(bank, 0, 1)
	}
	return sum
}

func solve2(data string) (sum int) {
	for _, bank := range strings.Split(data, "\n") {
		sum += findSubJolt(bank, 0, 11)
	}
	return sum
}

var memo = map[string]int{} // Not sure how long it would take without memoization, but likely forever

func findSubJolt(bank string, start int, depth int) int {
	key := fmt.Sprintf("%s|%d|%d", bank, start, depth)
	if v, ok := memo[key]; ok {
		return v
	}

	l := len(bank)
	if start == l-1 {
		last, _ := strconv.Atoi(bank[l-1:])
		return last
	}
	largest := 0
	factor := int(math.Pow10(depth))
	for d := start; d < (l - depth); d++ {
		cur, _ := strconv.Atoi(bank[d : d+1])
		if (cur+1)*factor-1 <= largest { // This alone makes the calculation 6 times faster
			continue
		}
		cur *= factor
		if depth != 0 {
			cur += findSubJolt(bank, d+1, depth-1)
		}
		largest = max(cur, largest)
	}

	memo[key] = largest
	return largest
}

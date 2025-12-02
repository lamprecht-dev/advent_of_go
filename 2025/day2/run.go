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

func getPatterns(d string) (int, int) {
	min, max, _ := strings.Cut(d, "-")
	minPattern := getMinPattern(min)
	maxPattern := getMaxPattern(max)

	return minPattern, maxPattern
}

func getMinPattern(num string) int {
	overflow, l, r := splitHalves(num)
	if overflow == 0 {
		if l < r {
			return l + 1
		}
		return l
	}
	return overflow + 1
}

func getMaxPattern(num string) int {
	overflow, l, r := splitHalves(num)
	if overflow == 0 {
		if l <= r {
			return l
		}
		return l - 1
	}
	return overflow
}

func splitHalves(num string) (int, int, int) {
	l := len(num)
	l2 := int(math.Floor(float64(l) / 2.0))
	right, _ := strconv.Atoi(num[l-l2:])
	left, _ := strconv.Atoi(num[l-(l2*2) : l-l2])
	if l2*2 < l {
		return int(math.Pow10(l2)) - 1, left, right // Just fills with 99999...  -  10**l2 if l2 = 3 then => 1000 - 1 => 999
	}
	return 0, left, right
}

func numLen(num int) int {
	return len(strconv.Itoa(num)) // There is propably a better way mathematically...
}

func solve1(data string) int {
	ids := make(map[int]struct{}) //  make a set, there might be overlap between numbers!
	for _, d := range strings.Split(data, ",") {
		min, max := getPatterns(d)
		println(d, min, max)
		for i := min; i <= max; i++ {
			// println(d, min, max, int(math.Pow10(numLen(i)))*i+i)
			num := int(math.Pow10(numLen(i)))*i + i
			ids[num] = struct{}{}
		}
	}
	sum := 0
	for v := range ids {
		sum += v
	}
	return sum
}

// not 29_940_925_364 too high

func solve2(data string) (sum int) {
	return 0
}

package aogutils

import (
	"strconv"
	"strings"
)

func IntSplit(input string, sep string) []int {
	splits := strings.Split(input, sep)
	ints := []int{}

	for _, s := range splits {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}

	return ints
}

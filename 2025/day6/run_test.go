package main

import (
	"aog/internal/aogutils"
	"testing"
)

var expect1 = 4277556
var expect2 = 3263827

func TestSolve1(t *testing.T) {
	v := solve1(aogutils.GetTest(1))

	if v != expect1 {
		t.Errorf("Part 1\nExpected: %v\nBut got: %v", expect1, v)
	}
}

func TestSolve2(t *testing.T) {
	v := solve2(aogutils.GetTest(1))

	if v != expect2 {
		t.Errorf("Part 2\nExpected: %v\nBut got: %v", expect2, v)
	}
}

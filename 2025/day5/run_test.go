package main

import (
	"aog/internal/aogutils"
	"testing"
)

var expect1 = 3
var expect2 = 14

func TestSolve(t *testing.T) {
	v, v2 := solve(aogutils.GetTest(1))

	if v != expect1 || v2 != expect2 {
		t.Errorf("Part 1\nExpected: %v\nBut got: %v", expect1, v)
	}
}

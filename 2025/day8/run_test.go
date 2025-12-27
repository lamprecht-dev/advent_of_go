package main

import (
	"aog/internal/aogutils"
	"testing"
)

var expect1 = 40
var expect2 = 25272

func TestSolve(t *testing.T) {
	v := solve(aogutils.GetTest(1), 10, 1)
	v2 := solve(aogutils.GetTest(1), 1000, 2)

	if v != expect1 {
		t.Errorf("Part 1\nExpected: %v\nBut got: %v", expect1, v)
	}
	if v2 != expect2 {
		t.Errorf("Part 2\nExpected: %v\nBut got: %v", expect2, v)
	}
}

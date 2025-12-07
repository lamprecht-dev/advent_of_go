package main

import (
	"aog/internal/aogutils"
	"testing"
)

var expect1 = 21
var expect2 = 40

func TestSolve(t *testing.T) {
	v, v2 := solve(aogutils.GetTest(1))

	if v != expect1 || v2 != expect2 {
		t.Errorf("\nExpected: %v, %v\nBut got: %v, %v", expect1, expect2, v, v2)
	}
}

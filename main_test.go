package main

import "testing"

func TestSolve(t *testing.T) {
	solutions := []int{0, 1, 0, 0, 2, 10, 4, 40, 92, 352, 724, 2680, 14200}
	for n, count := range solutions {
		if n > 1 && Solve(n) != count {
			t.Fail()
		}
	}
}

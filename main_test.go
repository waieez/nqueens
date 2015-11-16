package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	solutions := []int{0, 1, 0, 0, 2, 10, 4, 40, 92, 352, 724, 2680, 14200, 73712, 365596, 2279184, 14772512}
	for n, count := range solutions {
		if n > 1 {
			if sln := Solve(n); sln != count {
				fmt.Println(n, sln, count)
				t.Fail()
			}
		}
	}
}

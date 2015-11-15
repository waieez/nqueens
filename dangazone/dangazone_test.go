package dangazone

import "testing"

func TestToggle(t *testing.T) {
	dz := New(3)
	row, col := 1, 1
	dz.Toggle(row, col)
	if dz.row[row] != true ||
		dz.col[col] != true ||
		dz.ld[row-col-1] != true ||
		dz.rd[row+col+1] != true {
		t.Fail()
	}
}

func TestIsSafe_Clean(t *testing.T) {
	dz := New(3)
	if !dz.IsSafe(1, 1) {
		t.Fail()
	}
}

func TestIsSafe_Dirty(t *testing.T) {
	n := 3
	dz := New(n)
	dz.Toggle(1, 1)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if dz.IsSafe(row, col) {
				t.Fail()
			}
		}
	}
}

func TestIsSafe_Safe(t *testing.T) {
	n := 4
	dz := New(n)
	dz.Toggle(0, 0)
	if !dz.IsSafe(1, 2) {
		t.Fail()
	}
}

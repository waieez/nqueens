package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/waieez/nqueens/dangazone"
	"github.com/waieez/nqueens/queen"
)

func main() {
	count := Solve(16)
	fmt.Println("Solutions found:", count)
}

func SolveSeed(n, seed int) int {
	qs := queen.Stack{}
	dz := dangazone.New(n)
	row := 0
	col := seed

	// place queen in starting poisition
	qs.Push(row, col)
	// trigger the dangazones
	dz.Toggle(row, col)
	// set starting row/column to be 1, 0
	row = 1
	col = 0

	// for each starting point, try to find all solutions
	count := 0
	// while the row we're on is less than the number of rows
	for row < n {
		// while the column we're on is less than the number of columns
		for col < n {
			// if a queen can be placed
			if row != n && dz.IsSafe(row, col) {
				// trigger the horizontals/diagonals of the current position to be dangerous
				dz.Toggle(row, col)
				// put it in the current stack of placed queens
				qs.Push(row, col)
				// move to next row at this_row+1,0
				col = 0
				row++
				// if a queen can't be placed, increment the column (and try again)
			} else {
				col++
			}
		}
		// we've run out of columns for this row..

		// if the row is equal to the final row + 1
		if row == n {
			// we found a valid solution, increment the number of solutions by 1
			count++
		}

		// lets try again from the last queen placed
		// pop stack of queens to get position of the previous queen's row, col
		q := qs.Pop()
		row, col = q.Row, q.Col
		// unset the dangazone for the old queen's poisition
		dz.Toggle(row, col)
		// Exhausted all possibilities from this seed position
		if row == 0 && col == seed {
			return count
		}

		// set current position to previous queen's row, col+1
		col++
	}
	return count
}

// Solve Solves nQueens
func Solve(n int) int {
	count := 0
	mid := (n - 1) / 2
	even := n%2 == 0

	wg := sync.WaitGroup{}
	wg.Add(mid + 1)

	ch := make(chan int, mid+1)
	t := time.Now()
	for seed := 0; seed <= mid; seed++ {
		go (func(s int) {
			ct := 0
			sln := SolveSeed(n, s)
			if s < mid || even {
				ct += sln
			}
			ct += sln
			ch <- ct
			wg.Done()
		})(seed)
	}

	wg.Wait()
	close(ch)
	for c := range ch {
		count += c
	}
	fmt.Println(n, time.Since(t))
	return count
}

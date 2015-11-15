package main

import "fmt"

func main() {
	// for each starting point, try to find all solutions

	// place queen in starting poisition
	// trigger the dangazones
	// set starting row/column to be 1, 0

	// while the row we're on is less than the number of rows
		// #// if the row we're on is 0 row and the column isn't the seed column
		// #	// exhausted all possibilities for this seed, return count

		// while the column we're on is less than the number of columns
			// if a queen can be placed
				// put it in the current stack of placed queens
				// trigger the horizontals/diagonals of the current position to be dangerous
				// move to next row at this_row+1,0
			// if a queen can't be placed,
				// increment the column (and try again)

		// we've run out of columns for this row..
		// if the row is equal to the final row + 1
			// increment the number of solutions by 1
		// else if the row is the home row
			// return count

		// pop stack of queens to get position of the previous queen's row, col
		// unset the dangazone for the old queen's poisition
		// set current position to previous queen's row, col+1
}

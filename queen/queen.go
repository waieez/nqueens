package queen

// Queen is a queen
type Queen struct {
	Row, Col int
}

// Stack a Stack of Queens
type Stack struct {
	queens []Queen
}

// Push pushes onto the stack
func (s *Stack) Push(row, col int) {
	s.queens = append(s.queens, Queen{row, col})
}

// Pop removes the last queen from the stack
func (s *Stack) Pop() Queen {
	popped := s.queens[len(s.queens)-1]
	s.queens = s.queens[:len(s.queens)-1]
	return popped
}

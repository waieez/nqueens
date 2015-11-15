package dangazone

type Dangazone struct {
	rd  map[int]bool
	ld  map[int]bool
	row []bool
	col []bool
}

// New creates a new Dangazone
func New(n int) Dangazone {
	return Dangazone{
		rd:  map[int]bool{},
		ld:  map[int]bool{},
		row: make([]bool, n),
		col: make([]bool, n),
	}
}

// IsSafe checks if we're in a DANGAZONE~!
func (d *Dangazone) IsSafe(row, col int) bool {
	left := row - col - 1
	right := row + col + 1
	if d.row[row] || d.col[col] || d.ld[left] || d.rd[right] {
		return false
	}
	return true
}

// Toggle sets the dangazone on/off
func (d *Dangazone) Toggle(row, col int) {
	left := row - col - 1
	right := row + col + 1
	d.row[row], d.col[col], d.ld[left], d.rd[right] = !d.row[row], !d.col[col], !d.ld[left], !d.rd[right]
}

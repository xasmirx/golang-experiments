func seatsInTheater(nCols int, nRows int, col int, row int) int {
	x := nCols - col + 1
	y := nRows - row
	return x * y
}

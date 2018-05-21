func matrixElementsSum(matrix [][]int) int {

	sum := 0
	var invalidColumn []int

	for _, v := range matrix {
		for i, j := range v {

			found := false
			for _, v := range invalidColumn {
				if v == i {
					found = true
					break
				}
			}

			if found == true {
				continue
			}

			if j == 0 {
				invalidColumn = append(invalidColumn, i)
				continue
			}
			sum += j
		}
	}

	return sum
}

func sortByHeight(a []int) []int {

	var narray []int

	for _, v := range a {
		if v > -1 {
			narray = append(narray, v)
		}
	}

	sort.Sort(sort.IntSlice(narray))

	index := 0

	for k, v := range a {
		if v > -1 {
			a[k] = narray[index]
			index++
		}
	}

	return a
}

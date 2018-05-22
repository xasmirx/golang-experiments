func isLucky(n int) bool {
	conv := strconv.Itoa(n)
	l := len(conv)

	mid := (l / 2) - 1

	firstSum := 0
	secondSum := 0

	for i := 0; i < l; i++ {
		r := int(conv[i] - 48)
		if i <= mid {
			firstSum += r
		} else {
			secondSum += r
		}
	}

	return firstSum == secondSum
}
